const app = new Vue({
    el: '#post-secret',
    data: {
        errors: [],
        secret: null,
        expireAfterViews: null,
        expireAfter: null,
        serverApiBaseUrl: 'http://127.0.0.1:8088/v1',
        requestFailed: null,
    },
    methods: {
        checkForm: function (e) {
            this.errors = [];

            if (this.expireAfterViews > 2 ** 32) {
                this.errors.push('expireAfterViews can not be greater than 2^32');
            }

            if (this.expireAfter > 2 ** 32) {
                this.errors.push('expireAfter can not be greater than 2^32');
            }

            e.preventDefault();
            this.sendRequest();
        },
        sendRequest: function () {
            const secretForm = this;
            axios.post(secretForm.serverApiBaseUrl + '/secret', {
                firstName: 'Fred',
                lastName: 'Flintstone'
            })
                .then(function (response) {
                    console.log(response);
                })
                .catch(function (error) {
                    secretForm.requestFailed = "Error: " + error.message + ". Try again!";
                });
        }
    }
})
