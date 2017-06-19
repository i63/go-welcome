# go-welcome

```sh
oc new-app mysql -e MYSQL_ROOT_PASSWORD=password
export MYSQL=$(oc get pods -l app=mysql -o jsonpath={.items[0].metadata.name})
oc cp ./script.sql $MYSQL:/tmp/script.sql
oc rsh $MYSQL
mysql -h 127.0.0.1 -u root -p < /tmp/script.sql #inside container.

oc new-app https://github.com/i63/store-inventory --name=inventory
oc env dc inventory sql_db=store sql_host=mysql sql_user=root sql_password=password
```
