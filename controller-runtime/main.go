package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	ctrl "sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func main() {

	fmt.Println("prometheus setting")
	one := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "test_one",
		Help: "test metric for testing",
	})

	fmt.Println("metrics.Registry")
	err := metrics.Registry.Register(one)
	if err != nil {
		fmt.Println(err)
	}

	//var listener net.Listener
	//fmt.Println("listener")
	/*listener, err = metrics.NewListener(":8081")
	if err != nil {
		fmt.Println(err)
	}*/

	opts := ctrl.Options{
		MetricsBindAddress: ":8081",
		/*newMetricsListener: func(addr string) (net.Listener, error) {
			var err error
			listener, err = metrics.NewListener(addr)
			return listener, err
		},*/
	}
	fmt.Println("create NewManager")
	cfg, err := config.GetConfig()
	if err != nil {
		fmt.Println(err)
	}
	mgr, err := ctrl.New(cfg, opts)
	if err != nil {
		fmt.Println(err)
	}
	err = mgr.AddMetricsExtraHandler("/debug", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("Some debug info"))
		one.Inc()
	}))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("create ctx")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("start server")
	mgr.Start(ctx)

	//fmt.Sprintf("http://localhost:8081/debug")
}
