
```
+----------------------------------------------------------------+
|                                                                |
|                                Dag                             |
|                                                                |
+----------+------------------------+---------------------+------+
           |                        |                     |
           |                        |                     |
           |                        |                     |
           |                        |                     |
+----------v---------+   +----------v--------+   +--------v------+
|                    +--->  UtxoRepository   |   |   Validator   |
|   UnitRepository   |   +-------------------+   +--------------++
|                    +---------------------------+              |
+------+----------+--------------+   |           |              |
       |          |              |   |           |              |
       |          |              |   |           |              |
  +----v-----+  +-v-------+  +---v---v--+  +-----v----+         |
  |  dagdb   |  |  idxdb  |  |  utxodb  |  |  statedb <---------+
  +------+---+  +----+----+  +----+-----+  +-----+----+
         |           |            |              |
         |           |            |              |
         |           |            |              |
         |     +-----v------------v------+       |
         |     |                         |       |
         +---> |     ptndb.Database      | <-----+
               |                         |
               +-------------------------+
```