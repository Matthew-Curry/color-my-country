/* Initialization script for the database. */
load("./users.js")
load("./counties.js")

USER_COLLECTION = "users"
COUNTY_COLLECTION = "counties"

USER_INDEX = "userID"
COUNTY_INDEX = "properties.GEO_ID"

if (!db.getCollectionNames().includes(USER_COLLECTION)) {
    addCollection(USER_COLLECTION, users, USER_INDEX)
}

if (!db.getCollectionNames().includes(COUNTY_COLLECTION)) {
    addCollection(COUNTY_COLLECTION, counties, COUNTY_INDEX)
}

/**
 * Function to add a collection to the database with given seed data and column to index.
 * @param {string} collectionName - the name of the collection to add to the database
 * @param {Array} seedData - an array holding objects to add to the collection upon initialization
 * @param {string} indexColumn - the name of secondary index to add to the collection
 */
const addCollection = (collectionName, seedData, indexColumn) => {
    // create collection
    db.createCollection(collectionName)

    // add any seed data
    db[collectionName].insertMany(seedData)

    // add any index
    db[collectionName].createIndex({ indexColumn: 1})
    print(`Added collection ${collectionName}`)
}
