import { Component, OnInit } from '@angular/core';
import * as L from 'leaflet';

const TILE_LAYER_URL = "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png";
const OPEN_STREET_ATTRIBUTION = '&copy; <a href="http://www.openstreetmap.org/copyright>OpenStreetMap></a>"';
const MAX_ZOOM = 18;
const MIN_ZOOM = 1;
const CENTER_LATTITUDE = 39.8355;
const CENTER_LONGITUDE = -99.0909;
const MAP_ZOOM = 4;

@Component({
  selector: 'app-map',
  templateUrl: './map.component.html',
  styleUrls: ['./map.component.css']
})
export class MapComponent implements OnInit{

  private map: L.Map;
  private centroid: L.LatLngExpression = [CENTER_LATTITUDE, CENTER_LONGITUDE];

  constructor() {}

  ngOnInit(): void {
    this.initMap();
  }

  initMap(): void {
    // create map
    this.map = L.map('map', {
      center: this.centroid,
      zoom: MAP_ZOOM
    })

    // add tiles
    const tiles = L.tileLayer(TILE_LAYER_URL, {
      maxZoom: MAX_ZOOM,
      minZoom: MIN_ZOOM,
      attribution: OPEN_STREET_ATTRIBUTION
    })  

    tiles.addTo(this.map)

  }

  ngAfterViewChecked(): void {
    this.map.invalidateSize(true);
    //this.map.center = this.center;
  }

}
