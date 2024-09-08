```mermaid
graph TB
    subgraph AWS
        subgraph Region
        ALB(Application Load Balancer) --> VPN(Client VPN)
            subgraph VPC
                VPN -- TCP/443 --> Postgres
                VPN --> S3
                VPN --> EC2
            end
            ACM(ACM)
        end
    end
    Client(Client) --> Route53(Route 53)
    Route53 --> ALB