#!/bin/bash


eksctl create cluster \
#  -n, --name string               EKS cluster name (generated if unspecified, e.g. "ridiculous-outfit-1601403687")
       --name nova-eks-cluster \
#      --tags stringToString       Used to tag the AWS resources. List of comma separated KV pairs "k1=v1,k2=v2" (default [])
#  -r, --region string             AWS region
       --region us-west-2 \
#      --with-oidc                 Enable the IAM OIDC provider
#      --zones strings             (auto-select if unspecified)
#      --version string            Kubernetes version (valid options: 1.14, 1.15, 1.16, 1.17) (default "1.17")
#  -f, --config-file string        load configuration from a file (or stdin if set to '-')
#      --timeout duration          maximum waiting time for any long-running operation (default 25m0s)
#      --install-vpc-controllers   Install VPC controller that's required for Windows workloads
#      --managed                   Create EKS-managed nodegroup
       --managed \
#      --fargate                   Create a Fargate profile scheduling pods in the default and kube-system namespaces onto Fargate
       --fargate \
#Initial nodegroup flags:
#      --nodegroup-name string          name of the nodegroup (generated if unspecified, e.g. "ng-bab3535f")
       --nodegroup-name nova-node \
#      --without-nodegroup              if set, initial nodegroup will not be created
#  -t, --node-type string               node instance type (default "m5.large")
#  -N, --nodes int                      total number of nodes (for a static ASG) (default 2)
       --nodes 1 \
#  -m, --nodes-min int                  minimum nodes in ASG (default 2)
#  -M, --nodes-max int                  maximum nodes in ASG (default 2)
#      --node-volume-size int           node volume size in GB (default 80)
#      --node-volume-type string        node volume type (valid options: gp2, io1, sc1, st1) (default "gp2")
#      --max-pods-per-node int          maximum number of pods per node (set automatically if unspecified)
#      --ssh-access                     control SSH access for nodes. Uses ~/.ssh/id_rsa.pub as default key path if enabled
       --ssh-access
#      --ssh-public-key string          SSH public key to use for nodes (import from local path, or use existing EC2 key pair)
#      --node-ami string                Advanced use cases only. If 'ssm' is supplied (default) then eksctl will use SSM Parameter; if 'auto' is supplied then eksctl will automatically set the AMI based on version/region/instance type; if static is supplied (deprecated), then static AMIs will be used; if any other value is supplied it will override the AMI to use for the nodes. Use with extreme care.
#      --node-ami-family string         Advanced use cases only. If 'AmazonLinux2' is supplied (default), then eksctl will use the official AWS EKS AMIs (Amazon Linux 2); if 'Ubuntu1804' is supplied, then eksctl will use the official Canonical EKS AMIs (Ubuntu 18.04). (default "AmazonLinux2")
#  -P, --node-private-networking        whether to make nodegroup networking private
#      --node-security-groups strings   Attach additional security groups to nodes, so that it can be used to allow extra ingress/egress access from/to pods
#      --node-labels stringToString     Extra labels to add when registering the nodes in the nodegroup. List of comma separated KV pairs "k1=v1,k2=v2" (default [])
#      --node-zones strings             (inherited from the cluster if unspecified)
#      --instance-prefix string         Add a prefix value in front of the instance's name.
#      --instance-name string           Overrides the default instance's name. It can be used in combination with the instance-prefix flag.
#      --disable-pod-imds               Blocks IMDS requests from non host networking pods
#
#Cluster and nodegroup add-ons flags:
#      --install-neuron-plugin    Install Neuron plugin for Inferentia nodes (default true)
#      --asg-access               enable IAM policy for cluster-autoscaler
#      --external-dns-access      enable IAM policy for external-dns
#      --full-ecr-access          enable full access to ECR
#      --appmesh-access           enable full access to AppMesh
#      --appmesh-preview-access   enable full access to AppMesh Preview
#      --alb-ingress-access       enable full access for alb-ingress-controller
#
#VPC networking flags:
#      --vpc-cidr ipNet                 global CIDR to use for VPC (default 192.168.0.0/16)
#      --vpc-private-subnets strings    re-use private subnets of an existing VPC
#      --vpc-public-subnets strings     re-use public subnets of an existing VPC
#      --vpc-from-kops-cluster string   re-use VPC from a given kops cluster
#      --vpc-nat-mode string            VPC NAT mode, valid options: HighlyAvailable, Single, Disable (default "Single")
#
#AWS client flags:
#  -p, --profile string        AWS credentials profile to use (overrides the AWS_PROFILE environment variable)
#      --cfn-role-arn string   IAM role used by CloudFormation to call AWS API on your behalf
#
#Output kubeconfig flags:
#      --kubeconfig string               path to write kubeconfig (incompatible with --auto-kubeconfig) (default "/home/nova/.kube/config")
#      --authenticator-role-arn string   AWS IAM role to assume for authenticator
#      --set-kubeconfig-context          if true then current-context will be set in kubeconfig; if a context is already set then it will be overwritten (default true)
#      --auto-kubeconfig                 save kubeconfig file by cluster name, e.g. "/home/nova/.kube/eksctl/clusters/ridiculous-outfit-1601403687"
#      --write-kubeconfig                toggle writing of kubeconfig (default true)

