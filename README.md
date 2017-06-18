# go-welcome

```sh
oc new-app mysql -e MYSQL_ROOT_PASSWORD=password
export MYSQL=$(oc get pods -l app=mysql -o jsonpath={.items[0].metadata.name})
oc rsh $MYSQL 
oc new-app debianmaster/store-inventory
oc env dc inventory sql_db=test sql_host=mysql sql_user=root sql_password=password
```
