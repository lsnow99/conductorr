import router from "@/router";

/*
    AuthService
    Some helper methods to be mixed into all views of the application
*/

export const AuthService = {
    methods: {
        /*
            Checks for a successful authorized response, and if so sets jwt and
            optionally redirects. Returns whether jwt was deleted
        */
        checkAuthorizedToken(response, redirect = true) {
            if (response.status == 200) {
                localStorage.setItem("jwt", response.data);
                if (redirect) {
                    router.push({ path: "/" });
                }
                return true;
            }
            return false;
        },
        /*
            Delete jwt
        */
        removeJWT() {
            localStorage.removeItem("jwt");
        },
        /*
            Checks for an unauthorized 401 resposne, and if so deletes jwt and 
            optionally redirects. Returns whether jwt was deleted
        */
        checkUnauthorizedToken(error, redirect = true) {
            if (error.response && error.response.status == 401) {
                this.removeJWT();
                if (redirect) {
                    router.push({ path: "/auth", params: { expired: true } });
                }
                return true;
            }
            return false;
        },
        /*
            Log out the current user by deleting the jwt and optionally
            redirecting
        */
        logout(redirect = true) {
            this.removeJWT();
            if (redirect) {
                router.push({ path: "/" });
            }
        }
    }
};
