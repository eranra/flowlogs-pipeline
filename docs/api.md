
## Prometheus encode API
Following is the supported API format for prometheus encode:

<pre>
 prom:
         metrics: list of prometheus metric definitions, each includes:
                 name: the metric name
                 type: (enum) one of the following:
                     gauge: single numerical value that can arbitrarily go up and down
                     counter: monotonically increasing counter whose value can only increase
                     histogram: counts samples in configurable buckets
                 valuekey: entry key from which to resolve metric value
                 labels: labels to be associated with the metric
                 buckets: histogram buckets
         port: port number to expose "/metrics" endpoint
         prefix: prefix added to each metric name
         expirytime: seconds of no-flow to wait before deleting prometheus data item
</pre>
## Kafka encode API
Following is the supported API format for kafka encode:

<pre>
 kafka:
         addr: address of kafka server
         topic: kafka topic to write to
         balancer: (enum) one of the following:
             roundRobin: RoundRobin balancer
             leastBytes: LeastBytes balancer
             hash: Hash balancer
             crc32: Crc32 balancer
             murmur2: Murmur2 balancer
         writeTimeout: timeout (in seconds) for write operation performed by the Writer
         readTimeout: timeout (in seconds) for read operation performed by the Writer
         batchBytes: limit the maximum size of a request in bytes before being sent to a partition
         batchSize: limit on how many messages will be buffered before being sent to a partition
</pre>
## Ingest collector API
Following is the supported API format for the netflow collector:

<pre>
 collector:
         hostName: the hostname to listen on
         port: the port number to listen on
</pre>
## Ingest Kafka API
Following is the supported API format for the kafka ingest:

<pre>
 kafka:
         brokers: list of kafka broker addresses
         topic: kafka topic to listen on
         groupid: separate groupid for each consumer on specified topic
         groupbalancers: list of balancing strategies (range, roundRobin, rackAffinity)
         startoffset: FirstOffset (least recent - default) or LastOffset (most recent) offset available for a partition
         batchreadtimeout: how often (in milliseconds) to process input
</pre>
## Aws ingest API
Following is the supported API format for Aws flow entries:

<pre>
 aws:
         fields: list of aws flow log fields
</pre>
## Transform Generic API
Following is the supported API format for generic transformations:

<pre>
 generic:
         rules: list of transform rules, each includes:
                 input: entry input field
                 output: entry output field
</pre>
## Transform Network API
Following is the supported API format for network transformations:

<pre>
 network:
         rules: list of transform rules, each includes:
                 input: entry input field
                 output: entry output field
                 type: (enum) one of the following:
                     conn_tracking: set output field to value of parameters field only for new flows by matching template in input field
                     add_regex_if: add output field if input field satisfies regex pattern from parameters field
                     add_if: add output field if input field satisfies criteria from parameters field
                     add_subnet: add output subnet field from input field and prefix length from parameters field
                     add_location: add output location fields from input
                     add_service: add output network service field from input port and parameters protocol field
                     add_kubernetes: add output kubernetes fields from input
                 parameters: parameters specific to type
         kubeconfigpath: path to kubeconfig file (optional)
</pre>
## Write Loki API
Following is the supported API format for writing to loki:

<pre>
 loki:
         url: the address of an existing Loki service to push the flows to
         tenantID: identifies the tenant for the request
         batchWait: maximum amount of time to wait before sending a batch
         batchSize: maximum batch size (in bytes) of logs to accumulate before sending
         timeout: maximum time to wait for a server to respond to a request
         minBackoff: initial backoff time for client connection between retries
         maxBackoff: maximum backoff time for client connection between retries
         maxRetries: maximum number of retries for client connections
         labels: map of record fields to be used as labels
         staticLabels: map of common labels to set on each flow
         ignoreList: map of record fields to be removed from the record
         clientConfig: clientConfig
         timestampLabel: label to use for time indexing
         timestampScale: timestamp units scale (e.g. for UNIX = 1s)
</pre>