DROP TYPE IF EXISTS Delivery;
CREATE TYPE Delivery AS (
                            name_delivery varchar(50),
                            phone varchar(12),
                            zip varchar(8),
                            city varchar(50),
                            address varchar(100),
                            region varchar(100),
                            email varchar(50)
                        );


DROP TYPE IF EXISTS Payment;
CREATE TYPE Payment AS (
                           transaction_id varchar(100),
                           request_id varchar(100),
                           currency varchar(10),
                           provider_str varchar(50),
                           amount integer,
                           payment_dt integer,
                           bank varchar(100),
                           delivery_cost integer,
                           goods_total integer,
                           custom_fee integer
                       );


DROP TYPE IF EXISTS Items;
CREATE TYPE Items AS (
                         chrt_id integer,
                         track_number varchar(100),
                         price integer,
                         rid varchar(100),
                         name_item varchar(50),
                         sale integer,
                         size varchar(10),
                         total_price integer,
                         nm_id integer,
                         brand varchar(100),
                         status integer
                     );




CREATE TABLE IF NOT EXISTS Orders (
                                      order_uid varchar(50),
                                      track_number varchar(100),
                                      entry varchar(10),
                                      delivery Delivery,
                                      payment Payment,
                                      items Items[],
                                      locale_chr varchar(5),
                                      internal_signature varchar(5),
                                      customer_id varchar(50),
                                      delivery_service varchar(20),
                                      shardkey varchar(2),
                                      sm_id integer,
                                      data_created timestamp,
                                      oof_shard varchar(2)
);













