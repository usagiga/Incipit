<template>
  <v-app dark>
    <v-navigation-drawer
      v-model="drawer"
      :mini-variant="miniVariant"
      :clipped="clipped"
      fixed
      app
    >
      <v-list>
        <v-list-item
          v-for="(item, i) in items"
          :key="i"
          :to="item.to"
          router
          exact
        >
          <v-list-item-action>
            <v-icon>{{ item.icon }}</v-icon>
          </v-list-item-action>
          <v-list-item-content>
            <v-list-item-title v-text="item.title" />
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
    <v-app-bar
      :clipped-left="clipped"
      fixed
      app
    >
      <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
      <v-toolbar-title v-text="title" />
    </v-app-bar>
    <v-main>
      <v-container>
        <nuxt />
      </v-container>
    </v-main>

    <!-- Error Message Snackbar -->
    <v-snackbar :value="hasErrorMsg" timeout="-1">
      {{ errorMsg }}

      <template v-slot:action="{ attrs }">
        <v-btn color="error" text v-bind="attrs" @click="clearErrorMsg()">
          Close
        </v-btn>
      </template>
    </v-snackbar>
  </v-app>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

  @Component
export default class Default extends Vue {
    clipped = false
    drawer = false
    items = [
      {
        icon: 'mdi-link',
        title: 'Links',
        to: '/link'
      },
      {
        icon: 'mdi-account-supervisor',
        title: 'Admin users',
        to: '/admin'
      },
      {
        icon: 'mdi-logout',
        title: 'Logout',
        to: '/logout'
      }
    ]

    miniVariant = false
    title = 'Incipit'

    errorMsg: string = ''

    // noinspection JSUnusedGlobalSymbols
    mounted () {
      window.addEventListener('unhandledrejection', (event) => {
        this.errorMsg = event.reason.message
      })
      window.addEventListener('error', (event) => {
        this.errorMsg = event.error.message
      })
    }

    get hasErrorMsg (): boolean {
      return this.errorMsg.length !== 0
    }

    clearErrorMsg () {
      this.errorMsg = ''
    }
}
</script>
