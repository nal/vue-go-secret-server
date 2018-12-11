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
            data = Qs.stringify(
                {
                    secret: this.secret,
                    expireAfterViews: this.expireAfterViews,
                    expireAfter: this.expireAfter,
                }
            );

            axios.post(secretForm.serverApiBaseUrl + '/secret',
                data,
                {
                    headers: { "Content-Type": "application/x-www-form-urlencoded", "accept": "application/json" }
                }
            )
                .then(function (response) {
                    console.log(response);
                })
                .catch(function (error) {

                    if (error.response) {
                        secretForm.requestFailed = "Error: " + error.response.data.error_message + ". Try again!";
                        // console.log(error.response.data);
                        // console.log(error.response.status);
                        // console.log(error.response.headers);
                    } else {
                        // console.log('Error', error.message);
                    }

                });
        }
    }
})
