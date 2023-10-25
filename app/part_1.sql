-- Выборка всех уникальных eventType у которых более 1000 событий:
   SELECT eventtype, COUNT(*) AS event_count
   FROM events
   GROUP BY eventtype
   HAVING COUNT(*) > 1000;

-- Выборка событий, которые произошли в первый день каждого месяца:
   SELECT *
   FROM events
   WHERE EXTRACT(DAY FROM eventtime) = 1;

-- Выборка пользователей, которые совершили более 3 различных eventType:
   SELECT userid
   FROM events
   GROUP BY userid
   HAVING COUNT(DISTINCT eventtype) > 3;
