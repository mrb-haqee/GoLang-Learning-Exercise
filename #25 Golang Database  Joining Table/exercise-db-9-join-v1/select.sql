-- TODO: answer here
SELECT
    reports.id as id,
    students.fullname as fullname,
    students.class as class,
    students.status as status,
    reports.study as study,
    reports.score as score
FROM
    reports
    INNER JOIN students ON reports.student_id = students.id
WHERE
    reports.score < 70
    AND students.status = 'active'
ORDER BY
    reports.score ASC