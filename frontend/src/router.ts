import { createMemoryHistory, createRouter } from "vue-router";

import Home from "./pages/Home.vue";
import Signup from "./pages/Signup.vue";
import Login from "./pages/Login.vue";
import VerifyAccount from "./pages/VerifyAccount.vue";

const routes = [
    {path: "/", component: Home },
    {path:"/signup", component: Signup},
    {path:"/login", component: Login},
    {path:"/verify", component: VerifyAccount}
]

const router = createRouter({
    history: createMemoryHistory(),
    routes,
})


export default router