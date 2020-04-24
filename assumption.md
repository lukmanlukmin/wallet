
# Result and Assumption

**Stack Information**  
*1. Database & Migration* : Using postgreSQL and Attached migration data. use mail name as password to login.<br>
*2. Separation Concern* : Using go-module but with MVC design pattern, no additional/custom module attached<br>
*3. Test Coverage* : Still below 50% because limitation of working time.<br>
*4. Project Delivery* : Using docker as one place delivery container.<br>

**Available Resource**
```bash
Method | Path                   | Auth
-------------------------------------------------
POST   | /auth/login            | All
GET    | /auth/logout           | All with auth
GET    | /user                  | All with auth
PUT    | /user                  | All with auth
POST   | /user                  | admin with auth
GET    | /user/:id              | admin with auth
PUT    | /user/:id              | admin with auth
DELETE | /user/:id              | admin with auth
GET    | /users                 | admin with auth
POST   | /transaction/topup     | All with auth
POST   | /transaction/transfer  | All with auth

```

__Asumption__<br>
1. No option for data archiving, so table *balance_bank* & *balance_bank_history* will be removed and not implemented
2. User need having auth level, then need to create another table or just add enumeration. I choose to simply put enumeration on *user* table
3. Adding creating data date on *user_balance_history* to track when the transaction is created
4. Instead of using `feature-commit` i use branch to separate feature merging
5. balance resume resource end point from table *user_balance* & *user_balance_history* not created because no request to that kind of resource, but the service is already created