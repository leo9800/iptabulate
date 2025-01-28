# iptabulate

Pretty-printing `iptables` output.

### Usage

```sh
sudo iptables -L -v -n | /path/to/iptabulates
```

Sample output:
```
Chain INPUT (policy ACCEPT 31M packets, 137G bytes)
pkts bytes target prot opt in out source destination extra
---- ----- ------ ---- --- -- --- ------ ----------- -----

Chain FORWARD (policy ACCEPT 0 packets, 0 bytes)
pkts bytes                   target prot opt      in      out    source destination                       extra
---- ----- ------------------------ ---- --- ------- -------- --------- ----------- ---------------------------
778K 1783M              DOCKER-USER    0  --       *        * 0.0.0.0/0   0.0.0.0/0                            
778K 1783M DOCKER-ISOLATION-STAGE-1    0  --       *        * 0.0.0.0/0   0.0.0.0/0                            
525K 1768M                   ACCEPT    0  --       *  docker0 0.0.0.0/0   0.0.0.0/0 ctstate RELATED,ESTABLISHED
   0     0                   DOCKER    0  --       *  docker0 0.0.0.0/0   0.0.0.0/0                            
253K   14M                   ACCEPT    0  -- docker0 !docker0 0.0.0.0/0   0.0.0.0/0                            
   0     0                   ACCEPT    0  -- docker0  docker0 0.0.0.0/0   0.0.0.0/0                            

Chain OUTPUT (policy ACCEPT 24M packets, 46G bytes)
pkts bytes target prot opt in out source destination extra
---- ----- ------ ---- --- -- --- ------ ----------- -----

Chain DOCKER (1 references)
pkts bytes target prot opt in out source destination extra
---- ----- ------ ---- --- -- --- ------ ----------- -----

Chain DOCKER-ISOLATION-STAGE-1 (1 references)
pkts bytes                   target prot opt      in      out    source destination extra
---- ----- ------------------------ ---- --- ------- -------- --------- ----------- -----
253K   14M DOCKER-ISOLATION-STAGE-2    0  -- docker0 !docker0 0.0.0.0/0   0.0.0.0/0      
778K 1783M                   RETURN    0  --       *        * 0.0.0.0/0   0.0.0.0/0      

Chain DOCKER-ISOLATION-STAGE-2 (1 references)
pkts bytes target prot opt in     out    source destination extra
---- ----- ------ ---- --- -- ------- --------- ----------- -----
   0     0   DROP    0  --  * docker0 0.0.0.0/0   0.0.0.0/0      
253K   14M RETURN    0  --  *       * 0.0.0.0/0   0.0.0.0/0      

Chain DOCKER-USER (1 references)
 pkts bytes target prot opt in out    source destination extra
----- ----- ------ ---- --- -- --- --------- ----------- -----
1123K 2593M RETURN    0  --  *   * 0.0.0.0/0   0.0.0.0/0      

```

### Acknowledgement

[Darragh O'Reilly](https://github.com/djoreilly) implemented this in Python: [Pretty print iptables output. Align columns and strip out comments.](https://gist.github.com/djoreilly/57ef65723bc8ad7ecdb899c5b8aaca47).
