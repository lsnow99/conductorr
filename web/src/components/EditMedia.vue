<template>
    <header class="modal-card-header">
    <p class="modal-card-title">{{ media.title }}</p>
  </header>
  <section class="modal-card-content w-96 min-w-full">
    <o-field label="Profile">
    <o-select v-model="media.profile_id" placeholder="Profile">
      <option v-for="profileOption in profiles" :key="profileOption.id" :value="profileOption.id">{{profileOption.name}}</option>
    </o-select>
    </o-field>
  </section>
  <footer class="modal-card-footer">
    <o-button @click="$emit('close')">Cancel</o-button>
    <div>
      <o-button variant="primary" @click="save">Save</o-button>
    </div>
  </footer>
</template>

<script>
import APIUtil from '../util/APIUtil'
export default {
    data() {
      return {
        profiles: []
      }
    },
    props: {
        media: {
            type: Object,
            default: function() {
                return {}
            }
        }
    },
    emits: ['close', 'submit'],
    methods: {
      save() {
        APIUtil.updateMedia(this.media.id, this.media.profile_id).then(() => {
          this.$emit('close')
        })
      }
    },
    mounted() {
      APIUtil.getProfiles().then(profiles => {
        this.profiles = profiles
      })
    }
}
</script>