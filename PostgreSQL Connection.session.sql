CREATE VIEW v_high_salary_employee AS
SELECT
emp_id,
emp_firstname,
emp_lastname,
emp_salary
FROM
tbl_employee
WHERE
emp_salary > 50000;