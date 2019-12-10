import router from "@/router";
import { SnackbarProgrammatic as Snackbar } from "buefy";

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
                    router.push({ path: "/auth"});
                    Snackbar.open({
                        message: "Session expired",
                        type: "is-danger",
                        indefinite: false,
                        duration: 10000
                    });
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
                router.push({ path: "/auth" });
            }
        },
        /*
            Get axios configuration with authorization header
        */
        axiosAuthConfig() {
            return {
                headers: {
                   Authorization: "bearer " + localStorage.getItem("jwt")
                }
             }
        },
        /*
            Utility to check if there is a currently logged-in session
        */
       loggedIn() {
           return localStorage.getItem("jwt") != null
       }
    }
};
