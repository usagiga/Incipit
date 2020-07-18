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

    <v-pagination
      v-model="page"
      :length="length"
      total-visible="7"
      @input="movePage"
    />
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

  @Component
export default class AdminList extends Vue {
    page: number = 1
    pageParam: number = 1
    length: number = 10

    items = [
      { title: 'Test A' },
      { title: 'Test B' },
      { title: 'Test C' },
      { title: 'Test D' }
    ]

    movePage () {
      if (this.pageParam === this.page) {
        return
      }

      this.$router.push(`/admin/${this.page}/`)
    }

    mounted () {
      // Parse parameters
      this.pageParam = parseInt(this.$route.params.page)
      if (isNaN(this.pageParam)) {
        this.pageParam = 1
      }

      this.page = this.pageParam

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
