# Passenger Go Exporter
Passenger status monitoring agent for Prometheus.

You can be used as a sidecar container for Passenger applications.
Implemented in golang, keeps memory consumption below 100MB, and runs on CPU 0.05 core.

## Supported version

- golang: 1.22
- Passenger: 6.0.18, or later.

All other versions have not been tested.

## How to use

We expect to using this as a sidecar.<br>
A example is in [here](./test/kubernetes/).

### Sidecar Pattern

Container images can be pulled from [ghcr.io/rakutentech/passenger-go-exporter](https://github.com/orgs/rakutentech/packages/container/package/passenger-go-exporter).<br>
Passsenger Go Exporter must be shared the directory specified by PASSENGER_INSTANCE_REGISTRY_DIR with passenger application.

Passenger Application side defines `PASSENGER_INSTANCE_REGISTRY_DIR` environment varibles,
and specify emptyDir volume.
Passenger Go Exporter side defines similarly.

```yaml
    spec:
      volumes:
        - name: tmp
          emptyDir: {}
      containers:
        - name: example
          env:
            - name: PASSENGER_INSTANCE_REGISTRY_DIR
              value: /tmp/ruby
          volumeMounts:
            - mountPath: /tmp/ruby
              name: tmp
        - name: passenger-exporter
          image: ghcr.io/rakutentech/passenger-go-exporter:v1.4.2
          imagePullPolicy: IfNotPresent
          env:
            - name: PASSENGER_INSTANCE_REGISTRY_DIR
              value: /tmp/ruby
          volumeMounts:
            - mountPath: /tmp/ruby
              name: tmp
```

Passenger Go Exporter works with the following resources definitions.

```yaml
          resources:
            limits:
              cpu: 20m
              memory: 60Mi
            requests:
              cpu: 20m
              memory: 60Mi
```

Health check of Passenger Go Exporter is following contents.

```yaml
          livenessProbe:
            initialDelaySeconds: 10
            timeoutSeconds: 2
            httpGet:
              path: /health
              port: 9768
          readinessProbe:
            timeoutSeconds: 1
            httpGet:
              path: /health
              port: 9768
```

Default of Port number is 9768.
But this could be changed by args defined.

```yaml
          - args:
            - "-port"
            - "9149"
```

## Running Options

|Name|Description|Default|
|:---|:---|:---|
|-port|Listening port number|9768|

## Collect Metrics

|Name|Labels|Description|
|:---|:---|:---|
|passenger_go_wait_list_size|name|number of requests in queue in each application|
|passenger_go_process_count|name|number of current application processes in each application|
|passenger_go_process_real_memory_bytes|name, pid|memory usage of process in each PID|
|passenger_go_process_processed|name, pid|the number of requests handled by each process|

