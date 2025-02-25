select name, username, email from users where email = 'alice@example.com';


select u.username, s.type_id from users as u right join subscription as s on u.user_id = s.user_id where u.username = 'alicej';

select u.username, t.type_name, t.cpu_limit, t.memory_limit, t.namespace_limit from users as u right join subscription as s on u.user_id = s.user_id right join subscription_types as t on s.type_id = t.type_id where u.username = 'alicej';


select n.namespace, n.active from users as u right join namespace as n on u.username = n.username where u.username = 'alicej';


-- Get the subscription details of the user with the username 'charlieb'. If the user has an active subscription, return the subscription details. If the user does not have an active subscription, return the default subscription details.
-- The default subscription details are the details of the subscription type with type_id = 1.
-- The subscription details should include the user's name, subscription type name, CPU limit, memory limit, namespace limit, transaction date, and validity.


SELECT
    u.name,
    COALESCE(st.type_name, default_st.type_name) AS type_name,
    COALESCE(st.cpu_limit, default_st.cpu_limit) AS cpu_limit,
    COALESCE(st.memory_limit, default_st.memory_limit) AS memory_limit,
    COALESCE(st.namespace_limit, default_st.namespace_limit) AS namespace_limit,
    t.transaction_date,
    t.validity
FROM
    users u
LEFT JOIN
    subscription s ON u.user_id = s.user_id
LEFT JOIN
    "Transactions" t ON s.transaction_id = t.transaction_id AND t.validity >= CURRENT_DATE
LEFT JOIN
    subscription_types st ON s.type_id = st.type_id
JOIN
    subscription_types default_st ON default_st.type_id = 1 and u.username = 'charlieb'
ORDER BY
    s.type_id DESC NULLS LAST
LIMIT 1;


-- Get username, namespace and active namespaces

Select
	u.username,
	n.namespace,
	n.active
from
	users u
join
	namespace n on u.username = n.username and u.username = 'alicej';

-- To get active user namespace and limit

SELECT
    COALESCE(st.type_name, default_st.type_name) AS type_name,
    COALESCE(st.cpu_limit, default_st.cpu_limit) AS cpu_limit,
    COALESCE(st.memory_limit, default_st.memory_limit) AS memory_limit,
    COALESCE(st.namespace_limit, default_st.namespace_limit) AS namespace_limit,
    COUNT(n.namespace_id) AS namespace_count,
    t.transaction_date,
    t.validity
FROM
    users u
LEFT JOIN
    subscription s ON u.user_id = s.user_id
LEFT JOIN
    "Transactions" t ON s.transaction_id = t.transaction_id AND t.validity >= CURRENT_DATE
LEFT JOIN
    subscription_types st ON s.type_id = st.type_id
LEFT JOIN
    namespace n ON u.username = n.username AND n.active = TRUE -- Count only active namespaces
JOIN
    subscription_types default_st ON default_st.type_id = 1 and u.username = 'bobsmith'
GROUP BY
    u.name, st.type_name, st.cpu_limit, st.memory_limit, st.namespace_limit,
    default_st.type_name, default_st.cpu_limit, default_st.memory_limit, default_st.namespace_limit,
    t.transaction_date, t.validity
ORDER BY
    t.transaction_date DESC NULLS LAST;

-- To get all the namespaces

SELECT
    COALESCE(st.type_name, default_st.type_name) AS type_name,
    COALESCE(st.cpu_limit, default_st.cpu_limit) AS cpu_limit,
    COALESCE(st.memory_limit, default_st.memory_limit) AS memory_limit,
    COALESCE(st.namespace_limit, default_st.namespace_limit) AS namespace_limit,
    COUNT(n.namespace_id) AS namespace_count,
    t.transaction_date,
    t.validity
FROM
    users u
LEFT JOIN
    subscription s ON u.user_id = s.user_id
LEFT JOIN
    "Transactions" t ON s.transaction_id = t.transaction_id AND t.validity >= CURRENT_DATE
LEFT JOIN
    subscription_types st ON s.type_id = st.type_id
LEFT JOIN
    namespace n ON u.username = n.username -- Count only active namespaces
JOIN
    subscription_types default_st ON default_st.type_id = 1 and u.username = 'bobsmith'
GROUP BY
    u.name, st.type_name, st.cpu_limit, st.memory_limit, st.namespace_limit,
    default_st.type_name, default_st.cpu_limit, default_st.memory_limit, default_st.namespace_limit,
    t.transaction_date, t.validity
ORDER BY
    t.transaction_date DESC NULLS LAST;


