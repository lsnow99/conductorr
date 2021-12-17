export default {
    data() {
        return {
            lastButton: null
        }
    },
    methods: {
        restoreFocus() {
            if (this.lastButton) {
                this.lastButton.focus()
            }
        }
    }
}
