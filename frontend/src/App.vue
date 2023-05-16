<template>
    <div id="app">
        <h1>Simple Blockchain</h1>
        <input v-model="data" placeholder="Enter data for new block" />
        <button @click="addBlock">Add Block</button>
        <BlockChainBlock v-for="block in blockchain" :key="block.Hash" :block="block" />
    </div>
</template>

<script>
import axios from 'axios';
import BlockChainBlock from './components/BlockChainBlock.vue';

export default {
    components: {
        BlockChainBlock
    },
    data() {
        return {
            blockchain: [],
            data: ''
        };
    },
    created() {
        this.fetchBlockchain();
    },
    methods: {
        fetchBlockchain() {
            axios.get('http://localhost:8080/get')
                .then(response => {
                    if (this.validateBlockchain(response.data)) {
                        this.blockchain = response.data;
                    } else {
                        console.error('Invalid blockchain data:', response.data);
                    }
                })
                .catch(error => {
                    console.log(error);
                });
        },
        validateBlockchain(blockchain) {
            if (!Array.isArray(blockchain)) {
                return false;
            }
            for (let block of blockchain) {
                if (!this.validateBlock(block)) {
                    return false;
                }
            }
            return true;
        },
        validateBlock(block) {
            return 'Index' in block && 'Hash' in block && 'Data' in block && 'PrevHash' in block;
        },
        addBlock() {
            axios.get('http://localhost:8080/write?data=' + this.data)
                .then(response => {
                    this.blockchain = response.data;
                    this.data = '';
                })
                .catch(error => {
                    console.log(error);
                });
        }
    }
}
</script>

