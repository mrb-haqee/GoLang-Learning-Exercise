SELECT
    *
FROM
    reports
WHERE
    final_score < 70
    or absent > 5