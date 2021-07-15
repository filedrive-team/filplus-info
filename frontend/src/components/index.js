import Vue from 'vue';
import { Button, Table, TableColumn, Tooltip, Pagination, Input, Loading, Row, Col, Popover } from 'element-ui';
import enLocale from 'element-ui/lib/locale/lang/en'
import locale from 'element-ui/lib/locale'
import BaseTable from './table/BaseTable.vue'
import AdvancedTable from './table/AdvancedTable.vue'
import Search from './search'
import Clipboard from 'clipboard';
// 设置语言
locale.use(enLocale)
Vue.component(BaseTable.name, BaseTable)
Vue.component(AdvancedTable.name, AdvancedTable)
Vue.component(Search.name, Search)
Vue.use(Button)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Tooltip)
Vue.use(Pagination)
Vue.use(Input)
Vue.use(Col)
Vue.use(Row)
Vue.use(Popover)
Vue.use(Loading.directive);


Vue.prototype.$loading = Loading.service;
Vue.prototype.Clipboard = Clipboard;

