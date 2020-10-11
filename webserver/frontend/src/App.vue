<template>
    <div class="c-wrapper">
        <header class="c-header">
            <header class="navbar navbar-light">

            </header>
        </header>
        <div class="c-body">
            <main class="c-main">
                <!-- Main content here -->

                <router-view></router-view>

            </main>
        </div>
        <footer class="c-footer">
            <!-- Footer content here -->
        </footer>
    </div>
</template>

<script>


export default {
    data() {
        return {
            message: null,
            text: null
        }
    },
    mounted() {
    },
    methods: {
        greet() {
            this.$toast.add({severity: 'info', summary: 'Hello '  + this.text});
            this.message = 'Hello ' + this.text;
        }
    },

}
</script>

<style scoped>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}

.app-container {
  text-align: center;
}

body #app .p-button {
  margin-left: .2em;
}

form {
  margin-top: 2em;
}
</style>