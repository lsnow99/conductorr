<template>
    <section>

    </section>
</template>

<script>
export default {
    data() {
        return {
            indexers: []
        }
    },
    mounted() {
        this.indexers = [
            {
                type: 'torznab',
                name: '',
                priority: 1,
                url: '',
                apiKey: '',
                categories: []
            }
        ]
    }
}
</script>