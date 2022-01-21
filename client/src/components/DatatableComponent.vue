<template>
    <div class="">
        <v-container>
            <v-card
            > 
                <v-toolbar
                  dense
                  flat
                  class="body-2 font-weight-bold"
                  color="grey lighten-2"
                  >Filter</v-toolbar
                >

                <v-row class="ma-1">
                    <v-col cols="12">
                        <v-text-field
                           v-model="search"
                           append-icon="search"
                           label="Search"
                            single-line
                            hide-details
                         ></v-text-field>
                    </v-col>

                    <v-col cols="12">
                        <date-picker id="my_picker"
                            :date-input="{placeholder:'Date range'}"
                            language="en"
                            apply-button-label="Apply"
                            :show-helper-buttons="false"
                            :switch-button-initial="true"
                            :is-monday-first="true"
                            @date-applied="applyDateFilter"
                            @date-reset="applyDateFilter"
                            :calendar-date-input=calendarInput
                        />
                    </v-col>
                </v-row>

                <v-row class="ma-1">
                    <v-col cols="12">
                        <v-data-table
                          :page="page"
                          :pageCount="numberOfPages"
                          :headers="headers"
                          :items="orders"
                          :options.sync="options"
                          :server-items-length="totalOrders"
                          :items-per-page="5"
                          :loading="loading"
                          class="elevation-1"
                        >
                            <template v-slot:item.DeliveredAmount="{ item }">
                                <span
                                >
                                  {{ item.DeliveredAmount == 0 ? '-' : item.DeliveredAmount }}
                                </span>
                            </template>
                        </v-data-table>
                    </v-col>
                </v-row>
            </v-card>
        </v-container>
    </div>
</template>
<style scoped></style>
<script>
import axios from "axios";
import DatePicker, {CalendarDialog} from 'vue-time-date-range-picker/dist/vdprDatePicker'

export default {
    components: {
        DatePicker,
        // eslint-disable-next-line vue/no-unused-components
        CalendarDialog
      },
    name: "DatatableComponent",
    data() {
        return {
                calendarInput: {
                labelStarts: "Start Date",
                labelEnds: "End Date",
                format: "YYYY-MM-DD"
            },
            search: '',
            start: '',
            end: '',
            page: 1,
            totalOrders: 0,
            numberOfPages: 0,
            orders: [],
            loading: true,
            options: {},
            headers: [
                { text: "Order Name", value: "OrderName" },
                { text: "Customer Company", value: "CustomerCompany" },
                { text: "Customer Name", value: "CustomerName" },
                { text: "Order Date", value: "CreatedAt" },
                { text: "Delivered Amount", value: "DeliveredAmount" },
                { text: "Total Amount", value: "TotalAmount" },
            ],
        };
    },
    watch: {
        options: {
            handler() {
                this.readDataFromAPI();
            },
        },
        search(){
            this.readDataFromAPI()
        },
        deep: true,
    },
    methods: {
        applyDateFilter(start,end){
            this.start = start.toLocaleDateString('fr-CA');
            this.end = end.toLocaleDateString('fr-CA');
            this.readDataFromAPI()
        },
        readDataFromAPI() {
            this.loading = true;

            const { page, itemsPerPage, sortBy, sortDesc} = this.options;
            let pageNumber = page - 1;
            let search = this.search.trim().toLowerCase();
            let query = 'size=' + itemsPerPage + '&page=' + pageNumber;

            if(this.start.length) {
                query += '&start=' + this.start + '&end=' + this.end;
            }

            if(search.length) {
                query += '&filter=' + encodeURIComponent(search);
            }

            if(sortBy.length) {
                query += '&sort=' + sortBy[0];
            }

            if(sortDesc.length) {
                query += '&sort_order=' + sortDesc[0];
            }

            axios
            .get(
              "http://localhost:8080/api/orders/?" + query
            )
            .then((response) => {
                this.loading = false;
                if( response.data.length ) {
                    this.orders = response.data;
                    this.totalOrders = response.data[0].total;
                    this.numberOfPages = response.data[0].total/page;
                } else {
                    this.orders = [];
                    this.totalOrders = 0;
                    this.numberOfPages = 0;
                }
            }).catch(err => {
                console.log(err)
            });
        },
    },
    mounted() {
        this.readDataFromAPI();
    },
};
</script>
