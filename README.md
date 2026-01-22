# Part 1: System design & architecture

### 1. How would you structure this Laravel application as it continues to grow?

- For a large laravel project, i would like to split it into module by domain, it could be dependant packages that each team can be responsible of it or modules in folder structure

- Regarding where the logic should lives, it should be moved entirely from the controller class, and split it into services classes by its domain also i would to depend on DTOs a lot because in php we don't know what is been passed through arrays so it is better to use something like DTO

- For maintainability i would like to depend on separating boundaries between modules that we mentioned before using interfaces (Contract) to decrease decoupling as we could

### 2. What would you refactor first, and why?

- I would like to start with the most critical part in refactor or to be more precise, i would start with the most part that we need to add features for it and it is currently hard and have many non intendant issues

### 3. How would you approach performance and scalability?

- In this scenario i would to start using cache aggressively, because it will help a lot on reducing database hits until we make the system more stable and more memory efficient, also we could start moving long requests to queue using horizon for actions like sending emails, notifications and bulk actions

- I don't want to move forward on something like micro service trend, or like fully sharding for database and all of those critical decisions


# Part 2: Focused code example

- Please check the code on the `app` directory

# Part 3: Trade-offs & reflection

- I designed the Inventory service to handle single-product checks to keep the API clean and reusable. To avoid performance issues in the loop, I would implement a 'Repository' pattern inside the Inventory domain that fetches all required IDs in a single query before the loop begins. In a real world example (production) i would try to move part of the complex logic into queue like order processing for it's status and generating invoices

- If i have time maybe we could build something like Idempotency key to be send on each critical endpoint like payment endpoint to prevent double requests

- As mentioned before i will ignoring thinking on scaling to microservices

- Yes, i used AI to prepare the bold lines and high level thinking and just taking multiple opinions from the other, otherwise all of this docs had been written by me
