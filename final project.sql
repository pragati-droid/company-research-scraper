WITH duplicates AS (
    SELECT
        financial_id,
        ROW_NUMBER() OVER (
            PARTITION BY company_id, fiscal_year
            ORDER BY financial_id
        ) AS row_num
    FROM financials
)
DELETE FROM financials
WHERE financial_id IN (
    SELECT financial_id
    FROM duplicates
    WHERE row_num > 1
);
SELECT * FROM financials;