## functional requirement
1. must have
   1. login / logout / password reset 
   2. get balance
   3. personal information
   4. current tariff
2. nice to have
   1. additional information like balance history


## domains:
user {
   id
   login
   hashed pass
}
## user endpoints:
```
/login(login, pass) token
/logout(token)
/get_balance(token) balance
/personal_info(token) info
/current_tariff(token) info
```

MODELS:
tariff
settings(change)
personal_data