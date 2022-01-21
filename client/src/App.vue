<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <span class="mr-2">PackForm</span>
      </div>

    </v-app-bar>

    <v-main>
      <DatatableComponent/>
    </v-main>
  </v-app>
</template>

<script>
import DatatableComponent from './components/DatatableComponent';


export default {
  name: 'App',

  components: {
    DatatableComponent
  },

  data: () => ({
    //
  }),
};
</script>
