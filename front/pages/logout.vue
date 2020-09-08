<template>
  <div>
    <v-row>
      <v-col
        justify="center"
        align="center"
      >
        <v-progress-circular
          indeterminate
          size="320"
        >
          Logging out...
        </v-progress-circular>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
/* eslint-disable no-console */

import { Vue, Component } from 'nuxt-property-decorator'
import TokenStore from '~/utils/token-store'

  @Component({
    layout: 'blank'
  })
export default class Login extends Vue {
  mounted () {
    TokenStore.accessToken = ''
    TokenStore.refreshToken = ''

    this.$router.push('/')
  }

  head () {
    return {
      title: 'Logout'
    }
  }
}
</script>
