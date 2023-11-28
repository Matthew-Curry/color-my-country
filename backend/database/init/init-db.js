/* Initialization script for Color My Country database. */

const fs = require('fs');
const { MongoClient, Db } = require('mongodb');

const URL = 'mongodb://db:27017';
const DATABASE = 'color-my-country-db';
const CONFIG = "collections.json";

// entrypoint to start main synchronous thread.
setUpDatabase(URL, DATABASE, CONFIG).catch(console.dir);

/**
 * Constructor for CollectionConfig object. These objects define configuration for 
 * all collection setup tasks.
 * @param {string} collectionName Name of collection.
 * @param {string} indexField Name of field to index.
 * @param {string} dataFile File to load into collection.
 */
function CollectionConfig(collectionName, indexField, dataFile) {
    if (collectionName === null || collectionName === undefined) {
        throw new Error("collectionName cannot be null or undefined");
    }

    if (indexField === null || indexField === undefined) {
        throw new Error("indexField cannot be null or undefined");
    }

    if (dataFile === null || dataFile === undefined) {
        throw new Error("dataFile cannot be null or undefined");
    }

    this.collectionName = collectionName;
    this.indexField = indexField;
    this.dataFile = dataFile;
};

/**
 * Factory function to construct a CollectionConfig from a JSON.
 * @param {Object} object 
 * @returns {CollectionConfig}
 */
function createCollectionConfigFromJson(object) {
    const collectionConfig = new CollectionConfig(object.collectionName, 
                                                object.indexField, 
                                                object.dataFile
                                            );
    return collectionConfig;
}
    

/**
 * Runs the Database setup tasks.
 * @param {string} url Path to connect to MongoDB server.
 * @param {string} database Name of database to use. Will create the database if it doesn't exist,
 *                      otherwise will connect to existing one.
 * @param {string} configFile Path to collection configuration file. Will create the collection 
 *                      according to the config if it does not exist, else it will ignore it.
 */
async function setUpDatabase(url, database, configFile) {
    // read in config data as list of json objects.
    let jsonConfigArray = readJsonFile(configFile);

    // convert list of config jsons to list of CollectionConfigs.
    let collectionConfigsArray = []; 
    jsonConfigArray.forEach((jsonConfig) => {
        try {
            const collctionConfig = createCollectionConfigFromJson(jsonConfig);
            collectionConfigsArray.push(collctionConfig);
        } catch (error) {
            errorLogAndFail("Error reading in config, exiting:", error);
        }
    });

    // connect to database, setup each collection that doesn't already exist.
    const client = new MongoClient(url);
    try {
        await client.connect();
        console.log("Connected to MongoDB server.");

        const db = client.db(database);
        console.log(`Connected to database ${database}.`);

        for (const collctionConfig of collectionConfigsArray) {
            await createAndLoadCollection(db, collctionConfig);
        }

    } catch (error) {
        errorLogAndFail("Error loading data, exiting:", error);
    } finally {
        await client.close();
    }
}

/**
 * Helper method logs a given error message and fails the script.
 * @param {string} errorMsg 
 * @param {Error} error 
 */
function errorLogAndFail(errorMsg, error) {
    console.error(errorMsg, error.stack);
    process.exit(1);
}

/**
 * Helper method to read in JSON from file path.
 * @param {string} filePath 
 * @returns {Object} Object holding data at path.
 */
function readJsonFile(filePath) {
    // read in data from file path.
    let data;  
    try {
        data = fs.readFileSync(filePath, "utf8");
    } catch (error) {
        errorLogAndFail("Error reading config file:", error);
    }

    // parse data as JSON object.
    let jsonData;
    try {
        jsonData = JSON.parse(data);
    } catch (error) {
        errorLogAndFail("Error parsing JSON:", error);
    }

    return jsonData;
}

/**
 * Function to create a collection and load data from a file into it using a batching strategy.
 * Will return if collection already exists.
 * @param {Db} db Mongo DB object.
 * @param {CollectionConfig} collectionConfig The config used to setup the collection.
 * @returns {Promise} Promise indicating method ran with no errors.
 */
async function createAndLoadCollection(db, collectionConfig) {
    // exit if collection already exists.
    const collections = await db.listCollections({ name: collectionConfig.collectionName }).toArray();
    if (collections.length > 0) {
        return infoLogAndPromise(`Collection "${collectionConfig.collectionName}" exists. Skipping creation and data load.`);
    } 
    
    // create collection and unique index.
    const collection = db.collection(collectionConfig.collectionName);

    const indexOptions = {
        unique: true,
    };
          
    await collection.createIndex({ [collectionConfig.indexField]: 1 }, indexOptions);
    console.log(`Created collection "${collectionConfig.collectionName}" with index ${collectionConfig.indexField}.`);

    // read list of collections in as object and load at once.
    let collectionJsonList = readJsonFile(collectionConfig.dataFile);
    await collection.insertMany(collectionJsonList)

    return infoLogAndPromise(`Successfully loaded ${collectionJsonList.length} documents in collection ${collection}`);
}

/**
 * Helper method logs a given info message and returns promise with same message.
 * @param {string} infoMsg 
 * @returns {Promise} Resolved promise with given message. 
 */
function infoLogAndPromise(infoMsg) {
    console.log(infoMsg);
    return Promise.resolve(infoMsg);
}
