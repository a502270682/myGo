DROP EXTENSION pg_cron;
-- 创建cron插件
CREATE EXTENSION pg_cron;
GRANT USAGE ON SCHEMA cron TO analysis;
GRANT EXECUTE ON FUNCTION cron.schedule_in_database(text,text,text,text,text,boolean) TO analysis;
GRANT EXECUTE ON FUNCTION cron.unschedule(bigint) TO analysis;
RESET SESSION AUTHORIZATION;
SET SESSION AUTHORIZATION analysis;

SELECT cron.schedule_in_database('user_behavior', '30 0 * * *', '', 'analysis');
-- 终止id为5的任务
select  cron.unschedule(5);