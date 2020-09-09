<template>
  <div>
    <v-row>
      <v-col>
        <h1>Incipit</h1>

        <p>Incipit is URL shortener service for private use.</p>

        <v-btn text to="/login">
          Login
        </v-btn>
        <v-btn text href="https://github.com/usagiga/Incipit">
          <v-icon>
            mdi-github
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'
import IncipitApi from '~/utils/incipit-api'

  @Component({
    layout: 'blank'
  })
export default class Index extends Vue {
  // noinspection JSUnusedGlobalSymbols
  mounted () {
    IncipitApi(this.$router).isInstalled()
  }

  head () {
    return {
      title: 'Incipit',
      titleTemplate: ''
    }
  }
}
</script>
