<template>
<div class="container">
    <div style="display: flex; width: 100%; justify-content: center;">
      <img src="@/assets/default.svg" class="logo-main">
    </div>
        <div class="modal-card login-container">
            <header class="modal-card-head">
                <p class="modal-card-title">{{firstRun?'Signup':'Login'}}</p>
            </header>
            <section class="modal-card-body">
                <b-field label="Username">
                    <b-input
                        v-model="username"
                        placeholder="Your username"
                        required
                        @keyup.enter.native="authenticate">
                    </b-input>
                </b-field>

                <b-field label="Password">
                    <b-input
                        type="password"
                        v-model="password"
                        password-reveal
                        placeholder="Your password"
                        required
                        @keyup.enter.native="authenticate">
                    </b-input>
                </b-field>

            </section>
            <footer class="modal-card-foot" style="display: flex; justify-content: right">
                <b-button :loading="isAuthenticating" @click="authenticate" type="is-primary">{{ firstRun?'Create Account':'Login'}}</b-button>
            </footer>
        </div>
</div>
</template>

<script>
import axios from 'axios'
import { SnackbarProgrammatic as Snackbar } from 'buefy'

export default {
  name: 'authenticate',
  components: {
  },
  data() {
    return {
        firstRun: false,
        isAuthenticating: false,
        username: '',
        password: ''
    }
  },
  methods: {
    authenticate: function() {
        this.isAuthenticating = true
        let creds = {
            username: this.username,
            password: this.password,
        }
        axios.post("/auth/" + (this.firstRun?'signup':'login'), creds)
        .then(() => {
            // Redirect
        })
        .catch(() => {
            Snackbar.open({
                message: 'Error logging in',
                type: 'is-danger',
                indefinite: false,
                duration: 10000,
            })
        })
        .finally(() => {
            this.isAuthenticating = false;
        })
    }
  },
  mounted () {
    axios.get("/auth/firstRun")
    .then((response) => {
        this.firstRun = response.data.first_run
    })
    .catch(() => {
        Snackbar.open({
            message: 'Error communicating with API',
            type: 'is-danger',
            indefinite: false,
            duration: 10000,
        })
    })
    .finally(() => {
        this.isSaving = false
    })
  }
}
</script>

<style scoped>
.login-container {
  width: 350px; 
  border: 3px solid rgb(240, 240, 240); 
  border-radius: 4px;
  margin-top: -90px;
}
.logo-main {
  width: 400px;
  position: relative;
  margin-top: -80px;
}
</style>