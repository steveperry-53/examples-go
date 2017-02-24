# Finding Handlers for Kubernetes REST Paths

Given a Kubernetes REST path and an HTTP verb, find the handler.

For example, suppose you're given the verb GET along with this path:

    /api/v1/namespaces/default/pods/my-repset-26ecl

Note: my-repset-26ecl is the name of a Pod.

Try it out:

    kubectl proxy --port=8080
    curl localhost:8080/api/v1/namespaces/default/pods/my-repset-26ecl

In
[k8s.io/kubernetes/apiserver/pkg/endpoints/installer.go](https://github.com/kubernetes/apiserver/blob/51bebaffa01be9dc28195140da276c2f39a10cd4/pkg/endpoints/installer.go)
see the `APIInstaller` type:

    type APIInstaller struct {
	    group             *APIGroupVersion
	    prefix            string
	    minRequestTimeout time.Duration
	}

`APIInstaller` has several methods:

* Install
* NewWebservice
* getResourceKind
* restMapping
* registerResourceHandlers

Look at the `registerResourceHandlers` method:

    ...
    resource, subresource, err := splitSubresource(path)
    ...
    fqKindToRegister, err := a.getResourceKind(path, storage)
	...
	kind := fqKindToRegister.Kind
    ...

    // I'll take a guess that resource is my-repset-26ecl, and kind is Pod.

    switch action.Verb {
	  case "GET":
        var handler restful.RouteFunction
		...
        handler = handlers.GetResource(getter, exporter, reqScope)
        handler = metrics.InstrumentRouteFunc(action.Verb, resource, handler)
        ...
        route := ws.GET(action.Path).To(handler).
        Doc(doc).
        Param(ws.QueryParameter("pretty", "If 'true', then the output is pretty printed.")).
        Operation("read"+namespaced+kind+strings.Title(subresource)+operationSuffix).
        Produces(append(storageMeta.ProducesMIMETypes(action.Verb), mediaTypes...)...).
        Returns(http.StatusOK, "OK", versionedObject).
        Writes(versionedObject)
      ...
      addParams(route, action.Params)
      ws.Route(route)
      ...

Look at `handlers.GetResource` and `metric.InstrumentRouteFunc`.

[k8s.io/apiserver/pkg/endpoints/handlers/rest.go](https://github.com/kubernetes/kubernetes/tree/2bb1e7581544b9bd059eafe6ac29775332e5a1d6/staging/src/k8s.io/apiserver/pkg/endpoints/handlers/rest.go)

    func GetResource(r rest.Getter, e rest.Exporter, scope RequestScope) restful.RouteFunction {
	    return getResourceHandler(scope, ...
        ...

[k8s.io/apiserver/pkg/endpoints/metrics/metrics.go](https://github.com/kubernetes/kubernetes/blob/3e3133bc596b48d78f3d0e6825de5ec7d96517eb/staging/src/k8s.io/apiserver/pkg/endpoints/metrics/metrics.go)

    func InstrumentRouteFunc(verb, resource string, routeFunc restful.RouteFunction) restful.RouteFunction {

	    return restful.RouteFunction( func(request *restful.Request, response *restful.Response) {...}
        )
        ...

Above, `restful.RouteFunction` is a type conversion.

So I think we have finally found the handler.


    func(request *restful.Request, response *restful.Response) {

		now := time.Now()

		delegate := &responseWriterDelegator{ResponseWriter: response.ResponseWriter}

		_, cn := response.ResponseWriter.(http.CloseNotifier)
		_, fl := response.ResponseWriter.(http.Flusher)
		_, hj := response.ResponseWriter.(http.Hijacker)

		var rw http.ResponseWriter

		if cn && fl && hj {
			rw = &fancyResponseWriterDelegator{delegate}
		} else {
			rw = delegate
		}

		response.ResponseWriter = rw

		routeFunc(request, response)  // This creates the response.

		if verb == "LIST" && strings.ToLower(request.QueryParameter("watch")) == "true" {
			verb = "WATCH"
		}

		Monitor(&verb, &resource, utilnet.GetHTTPClient(request.Request), rw.Header().Get("Content-Type"), delegate.status, now)
	}



Look at the function that actually creates the response: routeFunc, also known as handlers.Resource.

    func GetResource(r rest.Getter, e rest.Exporter, scope RequestScope) restful.RouteFunction {

	  return getResourceHandler(scope,

		  func(ctx request.Context, name string, req *restful.Request) (runtime.Object, error) {
			// For performance tracking purposes.
			trace := utiltrace.New("Get " + req.Request.URL.Path)
			defer trace.LogIfLong(500 * time.Millisecond)

			// check for export
			options := metav1.GetOptions{}
			if values := req.Request.URL.Query(); len(values) > 0 {
				exports := metav1.ExportOptions{}
				if err := metainternalversion.ParameterCodec.DecodeParameters(values, scope.MetaGroupVersion, &exports); err != nil {
					return nil, err
				}
				if exports.Export {
					if e == nil {
						return nil, errors.NewBadRequest(fmt.Sprintf("export of %q is not supported", scope.Resource.Resource))
					}
					return e.Export(ctx, name, exports)
				}
				if err := metainternalversion.ParameterCodec.DecodeParameters(values, scope.MetaGroupVersion, &options); err != nil {
					return nil, err
				}
			}

			return r.Get(ctx, name, &options)
		})
    }

Wow, get used to the idea of passing functions around.

Look at `getResourceHandler`.

    func getResourceHandler(scope RequestScope, getter getterFunc) restful.RouteFunction {

	  return func(req *restful.Request, res *restful.Response) {

		w := res.ResponseWriter
		namespace, name, err := scope.Namer.Name(req)
		if err != nil {
			scope.err(err, res.ResponseWriter, req.Request)
			return
		}
		ctx := scope.ContextFunc(req)
		ctx = request.WithNamespace(ctx, namespace)

		result, err := getter(ctx, name, req)
		if err != nil {
			scope.err(err, res.ResponseWriter, req.Request)
			return
		}
		if err := setSelfLink(result, req, scope.Namer); err != nil {
			scope.err(err, res.ResponseWriter, req.Request)
			return
		}

		responsewriters.WriteObject(http.StatusOK, scope.Kind.GroupVersion(), scope.Serializer, result, w, req.Request)
	  }
    }

Finally, I think we have the function that writes the response.

## Notes

I'm just getting used to this idea of passing functions around.
I think this is what's happening.

A is a function that take a function and returns a function.
Suppose, for example, A returns a function that is better than the input function.

B is also a function that takes a function and returns a better function.

Suppose f is a function that does a certain job, but not very well.

Let g = fA. Or we could write g = A(f).

Now g does a pretty good job, but we want a function that does an even better job.

Let h = gB. Or we could write h = B(g).

Finally, we have the function h, which does the job really well.

h = fAB. Or we could write h = B(A(f)).

























