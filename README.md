# Golang-ObservabilityStack

label_values(container_cpu_usage_seconds_total,container_label_io_kubernetes_pod_namespace)


sum(
    rate(
        container_cpu_usage_seconds_total
        {
            container_label_io_kubernetes_pod_namespace=~"$namespace",
            image!=""
        }[1m]
    )
)
by (
    container_label_io_kubernetes_pod_name,
    container_label_io_kubernetes_container_name
) /
sum(
    container_spec_cpu_quota{
        container_label_io_kubernetes_pod_namespace=~"$namespace",
        image!=""
    }
    /
    container_spec_cpu_period{
        container_label_io_kubernetes_pod_namespace=~"$namespace",
        image!=""
    }
)
by(
    container_label_io_kubernetes_pod_name,
    container_label_io_kubernetes_container_name
)