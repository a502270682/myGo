BEGIN;
INSERT INTO school ( name, school_type )
VALUES
( '绵阳中学', 'h school' );

INSERT INTO school ( name, school_type )
VALUES
( '南山中学', 'j school' );

INSERT INTO user ( name, age, school_id )
VALUES
( 'gfy', 19, 1 );
COMMIT;