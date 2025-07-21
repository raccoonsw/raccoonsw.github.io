---
sidebar_position: 4
tags:
  - test
---

# Question 4

```text
Найдите ошибки в теле запроса. Документация https://iftfintech.testsbi.sberbank.ru:9443/fintech/api/v1/payments/from-invoice
```

**Ответ**

```diff
{
   "externalId": "d5ca0566-a7a2-423a-a0cb-1d599e56f231",
   "paymentNumber": "584",
+  "date": "2025-05-29",
-  "amount": 10,
+  "amount": "1.01", // у вас два разных типа указано на странице
-  "operationCode": 01,
+  "operationCode": "01",
+  "deliveryKind": "электронно",
   "priority": "5", // такие странные вещи не входили в задание
   "urgencyCode": "INTERNAL",
   "purpose": "Пополнение счета. НДС не облагается",
   "payeeAccount": "40702810806000002425",
   "vat": {
     "type": "ONTOP"
     "rate": "10",
     "amount": 2.0
   },
+  "linkedDocs": [
+    {
+      "docExtId": "31663ef5-7975-4016-b0f3-f1d70a4e9c22",
+      "type": "string"
+    }
+  ],
+  "payeeOrgIdHash": "a1b2c3d4e5f6a7b8c9d0e1f2a3b4c5d6e7f8a9b0c1d2e3f4a5b6c7d8e9f0a1b2",
+  "expirationDate": "2025-05-30",
   "orderNumber": "123"
}
```