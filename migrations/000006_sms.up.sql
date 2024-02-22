CREATE TABLE IF NOT EXISTS sms(
    id VARCHAR(255) PRIMARY KEY,
    phone VARCHAR(12) NOT NULL,
    code VARCHAR(4) NOT NULL,
    confirm BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);


CREATE OR REPLACE FUNCTION check_sms_code()
RETURNS TRIGGER AS
$$ 
BEGIN 
    IF NEW.confirm = TRUE AND now() - NEW.created_at > interval '1 minutes' THEN
        RAISE EXCEPTION 'You can not change confirm status';
    END IF;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER sms_confirm
BEFORE UPDATE ON "sms"
FOR EACH ROW
EXECUTE PROCEDURE check_sms_code();
