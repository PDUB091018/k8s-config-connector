// Copyright 2022 Google LLC. All Rights Reserved.
// 
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// 
//     http://www.apache.org/licenses/LICENSE-2.0
// 
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_node_pool blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containeraws/beta/node_pool.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/containeraws/beta/node_pool.yaml
var YAML_node_pool = []byte("info:\n  title: ContainerAws/NodePool\n  description: An Anthos node pool running on AWS.\n  x-dcl-struct-name: NodePool\n  x-dcl-has-iam: false\n  x-dcl-ref:\n    text: API reference\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest/v1/projects.locations.awsClusters.awsNodePools\n  x-dcl-guides:\n  - text: Multicloud overview\n    url: https://cloud.google.com/anthos/clusters/docs/multi-cloud\npaths:\n  get:\n    description: The function used to get information about a NodePool\n    parameters:\n    - name: NodePool\n      required: true\n      description: A full instance of a NodePool\n  apply:\n    description: The function used to apply information about a NodePool\n    parameters:\n    - name: NodePool\n      required: true\n      description: A full instance of a NodePool\n  delete:\n    description: The function used to delete a NodePool\n    parameters:\n    - name: NodePool\n      required: true\n      description: A full instance of a NodePool\n  deleteAll:\n    description: The function used to delete all NodePool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: cluster\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many NodePool\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n    - name: location\n      required: true\n      schema:\n        type: string\n    - name: cluster\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    NodePool:\n      title: NodePool\n      x-dcl-id: projects/{{project}}/locations/{{location}}/awsClusters/{{cluster}}/awsNodePools/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-has-iam: false\n      type: object\n      required:\n      - name\n      - version\n      - config\n      - autoscaling\n      - subnetId\n      - maxPodsConstraint\n      - project\n      - location\n      - cluster\n      properties:\n        annotations:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Annotations\n          description: 'Optional. Annotations on the node pool. This field has the\n            same restrictions as Kubernetes annotations. The total size of all keys\n            and values combined is limited to 256k. Key can have 2 segments: prefix\n            (optional) and name (required), separated by a slash (/). Prefix must\n            be a DNS subdomain. Name must be 63 characters or less, begin and end\n            with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics\n            between.'\n        autoscaling:\n          type: object\n          x-dcl-go-name: Autoscaling\n          x-dcl-go-type: NodePoolAutoscaling\n          description: Autoscaler configuration for this node pool.\n          required:\n          - minNodeCount\n          - maxNodeCount\n          properties:\n            maxNodeCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxNodeCount\n              description: Maximum number of nodes in the NodePool. Must be >= min_node_count.\n            minNodeCount:\n              type: integer\n              format: int64\n              x-dcl-go-name: MinNodeCount\n              description: Minimum number of nodes in the NodePool. Must be >= 1 and\n                <= max_node_count.\n        cluster:\n          type: string\n          x-dcl-go-name: Cluster\n          description: The awsCluster for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Gkemulticloud/Cluster\n            field: name\n            parent: true\n        config:\n          type: object\n          x-dcl-go-name: Config\n          x-dcl-go-type: NodePoolConfig\n          description: The configuration of the node pool.\n          required:\n          - iamInstanceProfile\n          - configEncryption\n          properties:\n            configEncryption:\n              type: object\n              x-dcl-go-name: ConfigEncryption\n              x-dcl-go-type: NodePoolConfigConfigEncryption\n              description: The ARN of the AWS KMS key used to encrypt node pool configuration.\n              required:\n              - kmsKeyArn\n              properties:\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: The ARN of the AWS KMS key used to encrypt node pool\n                    configuration.\n            iamInstanceProfile:\n              type: string\n              x-dcl-go-name: IamInstanceProfile\n              description: The name of the AWS IAM role assigned to nodes in the pool.\n              x-kubernetes-immutable: true\n            imageType:\n              type: string\n              x-dcl-go-name: ImageType\n              description: The OS image type to use on node pool instances.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n            instancePlacement:\n              type: object\n              x-dcl-go-name: InstancePlacement\n              x-dcl-go-type: NodePoolConfigInstancePlacement\n              description: Details of placement information for an instance.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              properties:\n                tenancy:\n                  type: string\n                  x-dcl-go-name: Tenancy\n                  x-dcl-go-type: NodePoolConfigInstancePlacementTenancyEnum\n                  description: 'The tenancy for the instance. Possible values: TENANCY_UNSPECIFIED,\n                    DEFAULT, DEDICATED, HOST'\n                  x-kubernetes-immutable: true\n                  x-dcl-server-default: true\n                  enum:\n                  - TENANCY_UNSPECIFIED\n                  - DEFAULT\n                  - DEDICATED\n                  - HOST\n            instanceType:\n              type: string\n              x-dcl-go-name: InstanceType\n              description: Optional. The AWS instance type. When unspecified, it defaults\n                to `m5.large`.\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n            labels:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Labels\n              description: 'Optional. The initial labels assigned to nodes of this\n                node pool. An object containing a list of \"key\": value pairs. Example:\n                { \"name\": \"wrench\", \"mass\": \"1.3kg\", \"count\": \"3\" }.'\n              x-kubernetes-immutable: true\n            proxyConfig:\n              type: object\n              x-dcl-go-name: ProxyConfig\n              x-dcl-go-type: NodePoolConfigProxyConfig\n              description: Proxy configuration for outbound HTTP(S) traffic.\n              required:\n              - secretArn\n              - secretVersion\n              properties:\n                secretArn:\n                  type: string\n                  x-dcl-go-name: SecretArn\n                  description: The ARN of the AWS Secret Manager secret that contains\n                    the HTTP(S) proxy configuration.\n                secretVersion:\n                  type: string\n                  x-dcl-go-name: SecretVersion\n                  description: The version string of the AWS Secret Manager secret\n                    that contains the HTTP(S) proxy configuration.\n            rootVolume:\n              type: object\n              x-dcl-go-name: RootVolume\n              x-dcl-go-type: NodePoolConfigRootVolume\n              description: Optional. Template for the root volume provisioned for\n                node pool nodes. Volumes will be provisioned in the availability zone\n                assigned to the node pool subnet. When unspecified, it defaults to\n                32 GiB with the GP2 volume type.\n              x-dcl-server-default: true\n              properties:\n                iops:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: Iops\n                  description: Optional. The number of I/O operations per second (IOPS)\n                    to provision for GP3 volume.\n                  x-dcl-server-default: true\n                kmsKeyArn:\n                  type: string\n                  x-dcl-go-name: KmsKeyArn\n                  description: Optional. The Amazon Resource Name (ARN) of the Customer\n                    Managed Key (CMK) used to encrypt AWS EBS volumes. If not specified,\n                    the default Amazon managed key associated to the AWS region where\n                    this cluster runs will be used.\n                sizeGib:\n                  type: integer\n                  format: int64\n                  x-dcl-go-name: SizeGib\n                  description: Optional. The size of the volume, in GiBs. When unspecified,\n                    a default value is provided. See the specific reference in the\n                    parent resource.\n                  x-dcl-server-default: true\n                volumeType:\n                  type: string\n                  x-dcl-go-name: VolumeType\n                  x-dcl-go-type: NodePoolConfigRootVolumeVolumeTypeEnum\n                  description: 'Optional. Type of the EBS volume. When unspecified,\n                    it defaults to GP2 volume. Possible values: VOLUME_TYPE_UNSPECIFIED,\n                    GP2, GP3'\n                  x-dcl-server-default: true\n                  enum:\n                  - VOLUME_TYPE_UNSPECIFIED\n                  - GP2\n                  - GP3\n            securityGroupIds:\n              type: array\n              x-dcl-go-name: SecurityGroupIds\n              description: Optional. The IDs of additional security groups to add\n                to nodes in this pool. The manager will automatically create security\n                groups with minimum rules needed for a functioning cluster.\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            sshConfig:\n              type: object\n              x-dcl-go-name: SshConfig\n              x-dcl-go-type: NodePoolConfigSshConfig\n              description: Optional. The SSH configuration.\n              required:\n              - ec2KeyPair\n              properties:\n                ec2KeyPair:\n                  type: string\n                  x-dcl-go-name: Ec2KeyPair\n                  description: The name of the EC2 key pair used to login into cluster\n                    machines.\n            tags:\n              type: object\n              additionalProperties:\n                type: string\n              x-dcl-go-name: Tags\n              description: Optional. Key/value metadata to assign to each underlying\n                AWS resource. Specify at most 50 pairs containing alphanumerics, spaces,\n                and symbols (.+-=_:@/). Keys can be up to 127 Unicode characters.\n                Values can be up to 255 Unicode characters.\n              x-kubernetes-immutable: true\n            taints:\n              type: array\n              x-dcl-go-name: Taints\n              description: Optional. The initial taints assigned to nodes of this\n                node pool.\n              x-kubernetes-immutable: true\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: object\n                x-dcl-go-type: NodePoolConfigTaints\n                required:\n                - key\n                - value\n                - effect\n                properties:\n                  effect:\n                    type: string\n                    x-dcl-go-name: Effect\n                    x-dcl-go-type: NodePoolConfigTaintsEffectEnum\n                    description: 'The taint effect. Possible values: EFFECT_UNSPECIFIED,\n                      NO_SCHEDULE, PREFER_NO_SCHEDULE, NO_EXECUTE'\n                    x-kubernetes-immutable: true\n                    enum:\n                    - EFFECT_UNSPECIFIED\n                    - NO_SCHEDULE\n                    - PREFER_NO_SCHEDULE\n                    - NO_EXECUTE\n                  key:\n                    type: string\n                    x-dcl-go-name: Key\n                    description: Key for the taint.\n                    x-kubernetes-immutable: true\n                  value:\n                    type: string\n                    x-dcl-go-name: Value\n                    description: Value for the taint.\n                    x-kubernetes-immutable: true\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: Output only. The time at which this node pool was created.\n          x-kubernetes-immutable: true\n        etag:\n          type: string\n          x-dcl-go-name: Etag\n          readOnly: true\n          description: Allows clients to perform consistent read-modify-writes through\n            optimistic concurrency control. May be sent on update and delete requests\n            to ensure the client has an up-to-date value before proceeding.\n          x-kubernetes-immutable: true\n        location:\n          type: string\n          x-dcl-go-name: Location\n          description: The location for the resource\n          x-kubernetes-immutable: true\n        maxPodsConstraint:\n          type: object\n          x-dcl-go-name: MaxPodsConstraint\n          x-dcl-go-type: NodePoolMaxPodsConstraint\n          description: The constraint on the maximum number of pods that can be run\n            simultaneously on a node in the node pool.\n          x-kubernetes-immutable: true\n          required:\n          - maxPodsPerNode\n          properties:\n            maxPodsPerNode:\n              type: integer\n              format: int64\n              x-dcl-go-name: MaxPodsPerNode\n              description: The maximum number of pods to schedule on a single node.\n              x-kubernetes-immutable: true\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The name of this resource.\n          x-kubernetes-immutable: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        reconciling:\n          type: boolean\n          x-dcl-go-name: Reconciling\n          readOnly: true\n          description: Output only. If set, there are currently changes in flight\n            to the node pool.\n          x-kubernetes-immutable: true\n        state:\n          type: string\n          x-dcl-go-name: State\n          x-dcl-go-type: NodePoolStateEnum\n          readOnly: true\n          description: 'Output only. The lifecycle state of the node pool. Possible\n            values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING,\n            ERROR, DEGRADED'\n          x-kubernetes-immutable: true\n          enum:\n          - STATE_UNSPECIFIED\n          - PROVISIONING\n          - RUNNING\n          - RECONCILING\n          - STOPPING\n          - ERROR\n          - DEGRADED\n        subnetId:\n          type: string\n          x-dcl-go-name: SubnetId\n          description: The subnet where the node pool node run.\n          x-kubernetes-immutable: true\n        uid:\n          type: string\n          x-dcl-go-name: Uid\n          readOnly: true\n          description: Output only. A globally unique identifier for the node pool.\n          x-kubernetes-immutable: true\n        updateTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: UpdateTime\n          readOnly: true\n          description: Output only. The time at which this node pool was last updated.\n          x-kubernetes-immutable: true\n        version:\n          type: string\n          x-dcl-go-name: Version\n          description: The Kubernetes version to run on this node pool (e.g. `1.19.10-gke.1000`).\n            You can list all supported versions on a given Google Cloud region by\n            calling GetAwsServerConfig.\n")

// 16338 bytes
// MD5: 2adcd193a525ba7f229e2235b01934e5
