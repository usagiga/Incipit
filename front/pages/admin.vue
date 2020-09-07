<template>
  <div class="admin-list">
    <v-list>
      <v-list-item
        v-for="item in items"
        :key="item.title"
      >
        <v-list-item-content>
          <v-list-item-title v-text="item.title" />
        </v-list-item-content>

        <v-list-item-action class="admin-list-action">
          <v-btn icon>
            <v-icon color="grey lighten-1">
              mdi-pencil
            </v-icon>
          </v-btn>
          <v-btn icon>
            <v-icon color="grey lighten-1">
              mdi-delete
            </v-icon>
          </v-btn>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

  @Component
export default class AdminList extends Vue {
    items = [
      { title: 'Test A' },
      { title: 'Test B' },
      { title: 'Test C' },
      { title: 'Test D' }
    ]

    mounted () {
      // TODO : Load admins
      // Temp
    }

    head () {
      return {
        title: 'Admins'
      }
    }
}
</script>

<style lang="scss">
  .admin-list {
    .admin-list-action {
      flex-direction: row;
    }
  }
</style>
