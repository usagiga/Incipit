<template>
  <div class="link-list">
    <v-list>
      <v-list-item>
        <v-list-item-content>
          <v-list-item-title>
            ID
          </v-list-item-title>
        </v-list-item-content>
        <v-list-item-content>
          <v-list-item-title>
            Actual URL
          </v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item
        v-for="(item, index) in linkItems"
        :key="item.title"
        class="link-list-item"
      >
        <v-list-item-content>
          <v-list-item-title v-show="!item.isEditing" v-text="item.id" />
          <v-list-item-title v-show="item.isEditing" v-text="item.id" />
        </v-list-item-content>

        <v-list-item-content>
          <v-list-item-subtitle v-show="!item.isEditing" v-text="item.actualUrl" />
          <v-text-field v-show="item.isEditing" v-model="item.editingUrl" placeholder="Actual URL" />
        </v-list-item-content>

        <v-list-item-action class="link-list-action">
          <v-btn v-show="!item.isEditing" icon @click="editLink(item)">
            <v-icon color="grey lighten-1">
              mdi-pencil
            </v-icon>
          </v-btn>
          <v-btn v-show="!item.isEditing" icon @click="deleteLink(item, index)">
            <v-icon color="grey lighten-1">
              mdi-delete
            </v-icon>
          </v-btn>
          <v-btn v-show="item.isEditing" icon @click="submitEditLink(item)">
            <v-icon color="grey lighten-1">
              mdi-checkbox-marked-circle-outline
            </v-icon>
          </v-btn>
          <v-btn v-show="item.isEditing" icon @click="cancelEditLink(item)">
            <v-icon color="grey lighten-1">
              mdi-close-circle-outline
            </v-icon>
          </v-btn>
        </v-list-item-action>
      </v-list-item>
    </v-list>
  </div>
</template>

<script lang="ts">
import { Vue, Component } from 'nuxt-property-decorator'

/**
 * An item of link list
 */
class LinkItem {
    id: string
    actualUrl: string
    editingUrl: string = ''
    isEditing: boolean = false

    constructor (id: string, actualUrl: string) {
      this.id = id
      this.actualUrl = actualUrl
    }
}

  @Component
export default class LinkList extends Vue {
    linkItems: LinkItem[] = [
      new LinkItem('Test A', 'https://example.com/'),
      new LinkItem('Test B', 'https://example.com/'),
      new LinkItem('Test C', 'https://example.com/')
    ]

    deleteLink (linkItem: LinkItem, index: number) {
      // TODO : Delete link through API
      // eslint-disable-next-line no-console
      console.log(`Deleted ${linkItem.id}` !)
      this.linkItems.splice(index, 1)
    }

    editLink (linkItem: LinkItem) {
      linkItem.editingUrl = linkItem.actualUrl

      // Change actions visibility
      linkItem.isEditing = true
    }

    submitEditLink (linkItem: LinkItem) {
      // TODO : Update link through API
      linkItem.actualUrl = linkItem.editingUrl

      // Change actions visibility
      linkItem.isEditing = false
    }

    cancelEditLink (linkItem: LinkItem) {
      // Change actions visibility
      linkItem.isEditing = false
    }

    mounted () {
      // TODO : Load links
      // Temp
    }

    head () {
      return {
        title: 'Links'
      }
    }
}
</script>

<style lang="scss">
  .link-list {
    .link-list-item {
      .link-list-action {
        flex-direction: row;
      }
    }
  }
</style>
