# cloud migration
### migrating app from on premise system to cloud from scratch.

everything begins from VPC. it is the very basic in the cloud infra. everything will be provisioned inside a Virtual Private Cloud. its your own set of machines. so lets first do that.

#### setting up the VPC

it will provide network isolation for our k8s cluster and associated resources. in creating vpc only, we have to take care of a lot of things so they don't cause a problem in the future. we never know what kind of setup we would have to go with. so creating a strong foundational block that is unbreakable in any situation should be the target. 
in setting up , we are choosing the 10.0.0.0/16 cidr block. we are not taking a very big one in case we need to 


















##### some points to consider
* we have to change all the defaut ports. for security purposes. *
* our application can be served via a public dns provided by aws. we can put it behind our actual domain name (to be done by ravi sir) and we are good to go. its not a problem. what we have to see is what all will be directly accessible to the customer/general public.*
* we also need to see if frontend needs multiple copies of pods or one is sufficient * 