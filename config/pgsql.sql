--
-- PostgreSQL database dump
--

-- Dumped from database version 16.1
-- Dumped by pg_dump version 16.1

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE IF EXISTS postgres;
--
-- Name: postgres; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE postgres WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Chinese (Simplified)_China.936';


ALTER DATABASE postgres OWNER TO postgres;

\connect postgres

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: DATABASE postgres; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE postgres IS 'default administrative connection database';


--
-- Name: react_go_manage; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA react_go_manage;


ALTER SCHEMA react_go_manage OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: department; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.department(
    _id character varying(100),
    dept_name character varying(50),
    user_name character varying(50),
    parent_id character varying(100) DEFAULT ''::character varying,
    create_id integer,
    update_time timestamp without time zone,
    create_time timestamp without time zone
);


ALTER TABLE react_go_manage.department OWNER TO postgres;

--
-- Name: line; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.line(
    label character varying(255),
    "order" integer,
    money integer
);


ALTER TABLE react_go_manage.line OWNER TO postgres;

--
-- Name: permission; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.permission(
    create_id integer,
    create_time timestamp with time zone,
    icon character varying(50),
    menu_name character varying(50),
    menu_state integer,
    menu_type integer,
    order_by integer,
    parent_id character varying(200),
    path character varying(20),
    update_time timestamp with time zone,
    _v integer DEFAULT 0,
    _id character varying(200),
    menu_code character varying(30)
);


ALTER TABLE react_go_manage.permission OWNER TO postgres;

--
-- Name: pie; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.pie(
    name character varying(255),
    value integer
);


ALTER TABLE react_go_manage.pie OWNER TO postgres;

--
-- Name: radar; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.radar(
    name character varying(255),
    online_duration integer,
    service_attitude integer,
    acceptance_rate integer,
    rating integer,
    popularity integer
);


ALTER TABLE react_go_manage.radar OWNER TO postgres;

--
-- Name: roles; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.roles(
    _id character varying(200),
    role_name character varying(100),
    checked_keys character varying(800),
    half_checked_keys character varying(800),
    create_time timestamp without time zone,
    update_time timestamp without time zone
);


ALTER TABLE react_go_manage.roles OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: react_go_manage; Owner: postgres
--

CREATE TABLE react_go_manage.users(
    create_id integer,
    create_time timestamp with time zone,
    dept_id character varying(100),
    dept_name character varying(100),
    job character varying(20),
    last_login_time timestamp without time zone,
    mobile character varying(25),
    role integer,
    role_list character varying(200),
    state integer,
    user_email character varying(50),
    user_id integer DEFAULT NOT NULL,
    user_img character varying(300),
    user_name character varying(50),
    _v integer,
    user_pwd character varying(50)
);


ALTER TABLE react_go_manage.users OWNER TO postgres;

--
-- Data for Name: department; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('655dbef811c02c8597dce77a', '大前端', 'CaiHongDi', '655dbeee11c02c8597dce776', 1000002, '2023-11-27 14:00:49', '2023-11-27 14:00:49');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('655dc07e11c02c8597dce7b5', '产品中心', 'CaiHongDi', '655dbeee11c02c8597dce776', 1000002, '2023-11-27 14:00:49', '2023-11-27 14:00:49');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('655dc06811c02c8597dce7ae', '测试部门', 'CaiHongDi', '655dbeee11c02c8597dce776', 1000002, '2023-11-27 14:00:49', '2023-11-27 14:00:49');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('655dc09311c02c8597dce7bd', '后端部门', 'CaiHongDi', '655dbeee11c02c8597dce776', 1000002, '2023-11-27 14:00:49', '2023-11-27 14:00:49');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('655dc08911c02c8597dce7b9', '产品中心', 'CaiHongDi', '655dbeee11c02c8597dce776', 1000002, '2023-11-27 14:00:49', '2023-11-27 14:00:49');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('655dbeee11c02c8597dce776', '技术中心', 'admin', '', 1000002, '2023-11-27 14:00:49', '2023-11-27 14:00:50');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('9c61d769-82c1-548a-93ac-41ca6ade7d68', '游戏发行部门', '疯驴子', '2d8bc600-1cc9-584a-8038-bab1fd35a5af', 0, '2023-11-27 15:37:59', '2023-11-27 15:37:59');
INSERT INTO react_go_manage.department (_id, dept_name, user_name, parent_id, create_id, update_time, create_time)VALUES ('2d8bc600-1cc9-584a-8038-bab1fd35a5af', '游戏中心', 'CaiHongDi', '', 0, '2023-11-27 15:38:09', '2023-11-27 15:37:13');


--
-- Data for Name: line; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.line (label, "order", money)VALUES ('5月', 463, 353);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('1月', 434, 43);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('12月', 391, 443);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('8月', 163, 446);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('2月', 217, 417);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('11月', 640, 67);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('6月', 958, 543);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('7月', 681, 361);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('4月', 228, 968);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('3月', 398, 701);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('9月', 27, 221);
INSERT INTO react_go_manage.line (label, "order", money)VALUES ('10月', 611, 252);


--
-- Data for Name: permission; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '查看', 1, 2, 0, '6083d756c30e1188761493f2', NULL, '2023-11-23 19:03:11.295+08', 0, '655db5a8f10762608048cb0a', 'menu@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '轨迹', 1, 2, 2, '6272009712eb226fad2f8e93', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc73fd4dc6d6fda15dbfa', 'order@route');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '查看', 1, 2, 0, '60979e4d3c0c8738d016ca60', NULL, '2023-11-23 19:03:11.295+08', 0, '655dbbc411c02c8597dce720', 'dept@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '删除', 1, 2, 3, '6069bec6b306e7f18dd72efd', NULL, '2023-11-23 18:46:04.694+08', 0, '655dc68ad4dc6d6fda15dbcb', 'user@delete');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '编辑', 1, 2, 2, '60979e4d3c0c8738d016ca60', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc6fdd4dc6d6fda15dbeb', 'dept@edit');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '新增', 1, 2, 1, '6083d756c30e1188761493f2', NULL, '2023-11-23 19:03:12.078+08', 0, '655dc69bd4dc6d6fda15dbcf', 'menu@create');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '新增', 1, 2, 1, '6083d76bc30e1188761493f3', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc6c7d4dc6d6fda15dbdb', 'role@create');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '删除', 1, 2, 3, '6272009712eb226fad2f8e93', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc74cd4dc6d6fda15dbfe', 'order@delete');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '新增', 1, 2, 1, '60979e4d3c0c8738d016ca60', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc6f4d4dc6d6fda15dbe7', 'dept@create');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 01:53:53.052+08', 'SettingOutlined', '系统管理', 1, 1, 2, '', '', '2023-11-21 01:55:33.186+08', 0, '6069beb8b306e7f18dd72efc', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:09:22.704+08', 'DatabaseOutlined', '订单管理', 1, 1, 3, '', '', '2023-11-21 02:09:57.13+08', 0, '6272005812eb226fad2f8e92', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '查看', 1, 2, 0, '63da0226a96e86702e4f2ca7', NULL, '2023-11-24 23:25:37+08', 0, '655dbc7911c02c8597dce734', 'cluster@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '查看', 1, 2, 0, '63e1ff8e65ac04da60e7a61c', NULL, '2023-11-25 00:13:03+08', 0, '655dbcb011c02c8597dce73c', 'driverList@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:00:42.017+08', 'MenuOutlined', '菜单管理', 1, 1, 2, '6069beb8b306e7f18dd72efc', '/menuList', '2023-11-25 03:02:32+08', 0, '6083d756c30e1188761493f2', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '编辑', 1, 2, 3, '6083d76bc30e1188761493f3', NULL, '2023-11-28 01:47:42+08', 0, '655dc6d3d4dc6d6fda15dbdf', 'role@edit');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '删除', 1, 2, 4, '6083d76bc30e1188761493f3', NULL, '2023-11-28 01:47:45+08', 0, '655dc6ddd4dc6d6fda15dbe3', 'role@delete');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:12:55.695+08', 'AimOutlined', '订单聚合', 1, 1, 2, '6272005812eb226fad2f8e92', '/cluster', '2023-11-21 02:12:53.251+08', 0, '63da0226a96e86702e4f2ca7', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '删除', 1, 2, 3, '6083d756c30e1188761493f2', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc6afd4dc6d6fda15dbd7', 'menu@delete');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:10:33.529+08', 'UnorderedListOutlined', '订单列表', 1, 1, 1, '6272005812eb226fad2f8e92', '/orderList', '2023-11-21 02:11:51.217+08', 0, '6272009712eb226fad2f8e93', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 01:56:16.569+08', 'TeamOutlined', '用户管理', 1, 1, 1, '6069beb8b306e7f18dd72efc', '/userList', '2023-11-21 01:57:17.852+08', 0, '6069bec6b306e7f18dd72efd', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:12:56.711+08', 'MehOutlined', '司机列表', 1, 1, 3, '6272005812eb226fad2f8e92', '/driverList', '2023-11-21 03:58:13.437+08', 0, '63e1ff8e65ac04da60e7a61c', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:01:51.697+08', 'TrademarkCircleOutlined', '角色管理', 1, 1, 3, '6069beb8b306e7f18dd72efc', '/roleList', '2023-11-21 02:02:37.969+08', 0, '6083d76bc30e1188761493f3', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-21 02:03:10.195+08', 'ProfileOutlined', '部门管理', 1, 1, 4, '6069beb8b306e7f18dd72efc', '/deptList', '2023-11-21 02:03:48.209+08', 0, '60979e4d3c0c8738d016ca60', NULL);
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '打点', 1, 2, 1, '6272009712eb226fad2f8e93', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc735d4dc6d6fda15dbf6', 'order@point');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '新增', 1, 2, 1, '6069bec6b306e7f18dd72efd', NULL, '2023-11-23 18:46:04.694+08', 0, '655dc53ed4dc6d6fda15dbad', 'user@create');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '编辑', 1, 2, 2, '6069bec6b306e7f18dd72efd', NULL, '2023-11-23 18:46:04.694+08', 0, '655dc67ed4dc6d6fda15dbc7', 'user@edit');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '删除', 1, 2, 3, '60979e4d3c0c8738d016ca60', NULL, '2023-11-23 19:03:11.295+08', 0, '"655dc708d4dc6d6fda15dbef', 'dept@delete');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '详情', 1, 2, 0, '6272009712eb226fad2f8e93', NULL, '2023-11-23 19:03:11.295+08', 0, '655dbc4d11c02c8597dce72c', 'order@detail');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '编辑', 1, 2, 2, '6083d756c30e1188761493f2', NULL, '2023-11-23 19:03:11.295+08', 0, '655dc6a5d4dc6d6fda15dbd3', 'menu@edit');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '查看', 1, 2, 0, '6083d76bc30e1188761493f3', NULL, '2023-11-23 19:03:11.295+08', 0, '655dbb0011c02c8597dce710', 'role@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2006-01-02 23:04:05+08', NULL, '查看', 1, 2, 0, '6069bec6b306e7f18dd72efd', NULL, '2023-11-23 18:46:04.694+08', 0, '655db556f10762608048cb02', 'user@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-25 03:01:07+08', '', '查看', 1, 2, 0, '623601eb-045b-51c6-b306-82fe888691f2', '', '2023-11-25 03:01:07+08', 0, 'e15d3cfc-de34-54c3-939a-bfaeacdca486', 'home@query');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-24 23:37:07+08', '', '新增', 1, 2, 0, '63f07a438c74bdc1580c2850', '', '2023-11-24 23:43:33+08', 0, 'f005a066-1ea3-5702-9684-f26a072166c1', 'home@create');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-25 03:00:08+08', 'DesktopOutlined', '工作台', 1, 1, 1, '', '/dashboard', '2023-11-25 03:00:22+08', 0, '623601eb-045b-51c6-b306-82fe888691f2', '');
INSERT INTO react_go_manage.permission (create_id, create_time, icon, menu_name, menu_state, menu_type, order_by, parent_id, path, update_time, _v, _id, menu_code)VALUES (1000002, '2023-11-28 01:45:12+08', '', '设置权限', 1, 2, 2, '6083d76bc30e1188761493f3', '', '2023-11-28 01:45:12+08', 0, '41b31ab8-c7d4-5f10-b6fb-85891fa43844', 'role@author');


--
-- Data for Name: pie; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.pie (name, value)VALUES ('北京', 123);
INSERT INTO react_go_manage.pie (name, value)VALUES ('上海', 456);
INSERT INTO react_go_manage.pie (name, value)VALUES ('深圳', 789);
INSERT INTO react_go_manage.pie (name, value)VALUES ('广州', 432);
INSERT INTO react_go_manage.pie (name, value)VALUES ('杭州', 765);


--
-- Data for Name: radar; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('李娜', 321, 142, 283, 288, 171);
INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('张三', 123, 234, 345, 456, 189);
INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('王五', 456, 397, 345, 128, 499);
INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('李四', 278, 178, 149, 432, 269);
INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('周宇', 192, 429, 305, 382, 127);
INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('王芳', 450, 462, 124, 215, 203);
INSERT INTO react_go_manage.radar (name, online_duration, service_attitude, acceptance_rate, rating, popularity)VALUES ('最大值', 600, 10, 100, 5, 10000);


--
-- Data for Name: roles; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.roles (_id, role_name, checked_keys, half_checked_keys, create_time, update_time)VALUES ('609781c15ccd183084f8ea3e', '产品经理', NULL, NULL, '2023-11-24 18:24:05', '2023-11-24 18:26:10');
INSERT INTO react_go_manage.roles (_id, role_name, checked_keys, half_checked_keys, create_time, update_time)VALUES ('63fe19eb03b115e52a6ac707', '市场部', NULL, NULL, '2023-11-24 18:26:09', '2023-11-24 18:26:08');
INSERT INTO react_go_manage.roles (_id, role_name, checked_keys, half_checked_keys, create_time, update_time)VALUES ('63bc3175300732c27697f1df', '研发', 'e15d3cfc-de34-54c3-939a-bfaeacdca486,655db556f10762608048cb02,655dc53ed4dc6d6fda15dbad,655dc67ed4dc6d6fda15dbc7,655dc68ad4dc6d6fda15dbcb,655db5a8f10762608048cb0a,655dc69bd4dc6d6fda15dbcf,655dc6a5d4dc6d6fda15dbd3,655dc6afd4dc6d6fda15dbd7,655dbb0011c02c8597dce710,655dc6c7d4dc6d6fda15dbdb,655dc6d3d4dc6d6fda15dbdf,655dc6ddd4dc6d6fda15dbe3,655dbbc411c02c8597dce720,655dc6f4d4dc6d6fda15dbe7,655dc6fdd4dc6d6fda15dbeb,"655dc708d4dc6d6fda15dbef,655dbc4d11c02c8597dce72c,655dc735d4dc6d6fda15dbf6,655dc73fd4dc6d6fda15dbfa,655dc74cd4dc6d6fda15dbfe,655dbc7911c02c8597dce734,655dbcb011c02c8597dce73c,41b31ab8-c7d4-5f10-b6fb-85891fa43844', '6069bec6b306e7f18dd72efd,6083d756c30e1188761493f2,60979e4d3c0c8738d016ca60,6272009712eb226fad2f8e93,63da0226a96e86702e4f2ca7,63e1ff8e65ac04da60e7a61c,623601eb-045b-51c6-b306-82fe888691f2,6272005812eb226fad2f8e92,6083d76bc30e1188761493f3,6069beb8b306e7f18dd72efc', '2023-11-24 18:26:10', '2023-11-27 17:48:59');
INSERT INTO react_go_manage.roles (_id, role_name, checked_keys, half_checked_keys, create_time, update_time)VALUES ('63fe19d503b115e52a6ac6fe', '研发经理', 'e15d3cfc-de34-54c3-939a-bfaeacdca486,655db556f10762608048cb02,655db5a8f10762608048cb0a,655dbb0011c02c8597dce710,655dbbc411c02c8597dce720,41b31ab8-c7d4-5f10-b6fb-85891fa43844', '623601eb-045b-51c6-b306-82fe888691f2,6069bec6b306e7f18dd72efd,6083d756c30e1188761493f2,6083d76bc30e1188761493f3,60979e4d3c0c8738d016ca60,6069beb8b306e7f18dd72efc', '2023-11-24 18:23:54', '2023-11-27 17:49:16');
INSERT INTO react_go_manage.roles (_id, role_name, checked_keys, half_checked_keys, create_time, update_time)VALUES ('63bc3187300732c27697f1e6', '测试', '655dbc4d11c02c8597dce72c,655dbc7911c02c8597dce734,655dbcb011c02c8597dce73c,e15d3cfc-de34-54c3-939a-bfaeacdca486', '63da0226a96e86702e4f2ca7,63e1ff8e65ac04da60e7a61c,623601eb-045b-51c6-b306-82fe888691f2,6272009712eb226fad2f8e93,6272005812eb226fad2f8e92', '2023-11-24 18:23:55', '2023-11-27 16:05:48');
INSERT INTO react_go_manage.roles (_id, role_name, checked_keys, half_checked_keys, create_time, update_time)VALUES ('63fe19f303b115e52a6ac70b', '运营部', NULL, NULL, '2023-11-24 18:23:51', '2023-11-24 18:23:53');


--
-- Data for Name: users; Type: TABLE DATA; Schema: react_go_manage; Owner: postgres
--

INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:19.048+08', '609781c15ccd183084f8ea3e', NULL, '产品经理', '2023-11-21 11:00:29.363', '18823456789', 3, '609781c15ccd183084f8ea3e', 1, 'demo1@example.com', 1000021, NULL, '徐江', 0, '1000021');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:18.27+08', '655dc07e11c02c8597dce7b5', NULL, '高级前端', '2023-11-21 11:00:30.252', '13567890123', 3, '63bc3175300732c27697f1df', 1, 'demo1@example.com', 1000020, NULL, '唐小虎', 0, '1000020');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:31:19.366+08', '63bc3175300732c27697f1df', NULL, '前端实习生1', '2023-11-21 11:32:03.207', '18934567890', 3, '63bc3175300732c27697f1df', 1, 'demo1@example.com', 1000026, NULL, '安欣', 0, '1000026');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:21.36+08', '9c61d769-82c1-548a-93ac-41ca6ade7d68', NULL, '市场部经理', '2023-11-21 11:00:31.151', '18934567890', 3, '63bc3175300732c27697f1df', 1, 'demo1@example.com', 1000023, NULL, '高启盛', 0, '1000023');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:16.536+08', '655dbef811c02c8597dce77a', '', '资深前端', '2023-11-21 11:00:28.54', '13678901234', 3, '63fe19d503b115e52a6ac6fe', 1, 'demo1@example.com', 1000018, NULL, '老默', 0, '1000018');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:15.515+08', '655dbef811c02c8597dce77a', '', '前端工程师', '2023-11-21 11:00:26.495', '18610996666', 2, '63bc3175300732c27697f1df', 1, '2696294007@qq.com', 1000016, NULL, 'CaiHongDi', 0, '123456');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:19.829+08', '2d8bc600-1cc9-584a-8038-bab1fd35a5af', '', '前端专家', '2023-11-21 11:00:35.742', '18798765432', 3, '63bc3175300732c27697f1df', 1, 'demo1@example.com', 1000022, NULL, '高启强', 0, '1000022');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:17.434+08', '655dbef811c02c8597dce77a', NULL, '产品经理', '2023-11-21 11:00:27.465', '13701234567', 3, '63bc3187300732c27697f1e6', 1, 'demo1@example.com', 1000019, NULL, '疯驴子', 0, '1000019');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:13.905+08', '655dbeee11c02c8597dce776', '', '全栈工程师', '2023-11-21 11:00:25.389', '18942909665', 0, '63fe19d503b115e52a6ac6fe', 1, '2250105638@qq.com', 1000002, NULL, 'admin', 0, '123456');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 22:31:42.728+08', '63bc3175300732c27697f1df', NULL, '后端go工程师', '2023-11-21 14:32:30.637', '13812345678', 1, NULL, 3, '2162797195@mars.com', 1000095, NULL, '王伟杰', 0, '2162797195');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 22:41:14.11+08', '63bc3175300732c27697f1df', NULL, '后端java实习生', '2023-11-21 14:41:15.861', NULL, 1, NULL, 2, '1628258955@mars.com', 1000051, NULL, '李雪婷', 0, '1000051');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:22.174+08', '63bc3187300732c27697f1e6', NULL, '测试专家', '2023-11-21 11:00:34.494', '15087654321', 3, '63bc3187300732c27697f1e6', 1, 'demo1@example.com', 1000024, NULL, '陈舒婷', 0, '1000024');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 22:41:14.887+08', '63bc3175300732c27697f1df', NULL, '前端vue实习生', '2023-11-21 14:41:16.835', NULL, 1, NULL, 2, '3425224885@mars.com', 1000084, NULL, '张晓琳', 0, '3425224885');
INSERT INTO react_go_manage.users (create_id, create_time, dept_id, dept_name, job, last_login_time, mobile, role, role_list, state, user_email, user_id, user_img, user_name, _v, user_pwd)VALUES (1000002, '2023-11-21 19:00:22.998+08', '63bc3187300732c27697f1e6', NULL, '测试专家', '2023-11-21 11:00:33.495', '15123456789', 3, '63bc3187300732c27697f1e6', 1, 'demo1@example.com', 1000025, NULL, '太叔', 0, '1000025');


--
-- PostgreSQL database dump complete
--

