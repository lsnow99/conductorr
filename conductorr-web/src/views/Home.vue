<template>
<div class="container">
    <div style="display: flex; width: 100%; justify-content: center;">
      <img src="@/assets/default.svg" class="logo-main">
    </div>
    <form action="">
        <div class="modal-card login-container">
            <header class="modal-card-head">
                <p class="modal-card-title">Login</p>
            </header>
            <section class="modal-card-body">
                <b-field label="Email">
                    <b-input
                        type="email"
                        :value="email"
                        placeholder="Your email"
                        required>
                    </b-input>
                </b-field>

                <b-field label="Password">
                    <b-input
                        type="password"
                        :value="password"
                        password-reveal
                        placeholder="Your password"
                        required>
                    </b-input>
                </b-field>

            </section>
            <footer class="modal-card-foot">
                <button class="button" type="button" @click="$parent.close()">Close</button>
                <button class="button is-primary">Login</button>
            </footer>
        </div>
    </form>
</div>
</template>

<script>

export default {
  name: 'home',
  components: {
  },
  props: {
    email: String,
    password: String
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