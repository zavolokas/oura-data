We do monitor different systems. Why not to monitor a human body as a system?

This project is to gather some metrics and represent them with Grafana. 

In the beginning there will be only Oura Ring as a source of the data.

```mermaid
flowchart LR
T[Telegraf] -->|input| B[body monitoring service]
B-->AR[Activity Reporter]
OAR[Oura Reporter]
C1AR[Client 1 Reporter]
C2AR[Client ... Reporter]
CNAR[Client N Reporter]
AR--> OAR
OAR-->OC[Oura Client] 
AR--> C1AR
C1AR-->C1[Client 1]
AR--> C2AR
C2AR-->C2[..........]
AR--> CNAR
CNAR-->CN[Client N]
T -->|output| IDB[InfluxDB]
G[Grafana]-->IDB

```
