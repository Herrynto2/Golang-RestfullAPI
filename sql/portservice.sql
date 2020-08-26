-- Create users Table
CREATE TABLE users (
    _id SERIAL PRIMARY KEY,
    email VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(256) NOT NULL,
	uuid VARCHAR(256),
    verify_code VARCHAR(256),
	created_at TIMESTAMP NOT NULL,
    modified_at TIMESTAMP NOT NULL,
    active_at TIMESTAMP NOT NULL,
    is_active BOOLEAN DEFAULT FALSE,
    is_verified BOOLEAN DEFAULT FALSE
);

CREATE TABLE client (
     _id SERIAL PRIMARY KEY,
     company_code INT NOT NULL,
     sales_org VARCHAR(256) NOT NULL,
     dist_channel VARCHAR(256) NOT NULL,
     division_client VARCHAR(256) NOT NULL,
     sap_customer_number INT,
     customer_name VARCHAR(256) UNIQUE NOT NULL,
     address VARCHAR(256) NOT NULL,
     telephone INT,
     mobile_phone INT,
     fax INT,
     email VARCHAR(256) NOT NULL,
     status VARCHAR(256) NOT NULL,
     attachment VARCHAR(256) NOT NULL,
     created_at TIMESTAMP NOT NULL,
     created_by VARCHAR(256) NOT NULL,
     modified_at TIMESTAMP,
     modified_by VARCHAR(256),
     delete_at TIMESTAMP,
     delete_by VARCHAR(256),
     user_id INT NOT NULL,
     FOREIGN KEY (user_id) REFERENCES users(_id) ON DELETE CASCADE
)

-- -- Create services table	
-- CREATE TABLE services (
-- 	_id SERIAL PRIMARY KEY,
-- 	request_id INT UNIQUE NOT NULL,
-- 	status VARCHAR(64) NOT NULL,
-- 	vessel_name VARCHAR(256) NOT NULL,
-- 	service_type VARCHAR(256) NOT NULL,
-- 	data_agent VARCHAR(256) NOT NULL,
-- 	cargo VARCHAR(256) NOT NULL,
-- 	etd VARCHAR(256) NOT NULL,
-- 	eta VARCHAR(256) NOT NULL,
--     user_id INT NOT NULL,
--     FOREIGN KEY (user_id) REFERENCES users(_id) ON DELETE CASCADE
-- );