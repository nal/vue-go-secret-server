const addSecret = new Vue({
    el: '#addSecret',
    data: {
        errors: [],
        secret: null,
        expireAfterViews: null,
        expireAfter: null,
        serverApiBaseUrl: 'http://127.0.0.1:8088/v1',
        requestFailed: null,
        requestSuccessful: { '_populated': 0 },
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
                    // console.log(response.data);
                    secretForm.requestSuccessful = response.data;
                    secretForm.requestSuccessful['_populated'] = 1;
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

const getSecret = new Vue({
    el: '#getSecret',
    data: {
        hash: null,
        serverApiBaseUrl: 'http://127.0.0.1:8088/v1',
        requestFailed: null,
        requestSuccessful: { '_populated': 0 },
    },
    methods: {
        sendRequest: function (e) {
            e.preventDefault();
            const getSecretForm = this;

            // Clear previous result if any
            getSecretForm.requestSuccessful['_populated'] = 0;
            getSecretForm.requestFailed = null;

            axios.get(getSecretForm.serverApiBaseUrl + '/secret/' + getSecretForm.hash,
                {
                    headers: { "Content-Type": "application/x-www-form-urlencoded", "accept": "application/json" }
                }
            )
                .then(function (response) {
                    getSecretForm.requestSuccessful = response.data;
                    getSecretForm.requestSuccessful['_populated'] = 1;
                })
                .catch(function (error) {

                    if (error.response) {
                        if (error.response.status == 404) {
                            // hash not found or expired
                            getSecretForm.requestFailed = "Secret with provided hash not found or already expired!"

                        }
                        else {
                            getSecretForm.requestFailed = "Error: " + error.response.data.error_message + ". Try again!";
                        }
                    } else {
                        // console.log('Error', error.message);
                    }
                });
        }
    }
})
