CREATE TABLE IF NOT EXISTS invoices (
    id TEXT PRIMARY KEY,
    total REAL NOT NULL,
    currency TEXT DEFAULT 'AOA',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS invoice_breakdown_steps (
    id TEXT PRIMARY KEY,
    invoice_id TEXT,
    step_order INTEGER NOT NULL,
    label TEXT NOT NULL,
    expression TEXT,
    formula TEXT NOT NULL,
    value REAL NOT NULL,
    description TEXT,
    FOREIGN KEY (invoice_id) REFERENCES invoices(id) ON DELETE CASCADE
);