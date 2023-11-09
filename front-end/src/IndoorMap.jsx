import React, { useEffect } from 'react';
import L from 'leaflet';
import 'leaflet/dist/leaflet.css';

function IndoorMap() {
	useEffect(() => {
		let footprintStyle = {
			color: 'black',
			fillColor: 'white',
			fillOpacity: 1
		};

		let L1footprint = L.polygon([
			[0, 0],
			[50.5, 0],
			[50.5, 40],
			[31.5, 40],
			[31.5, 35],
			[12, 35],
			[12, 17],
			[8, 17],
			[4, 17],
			[4, 10],
			[0, 10]
		], footprintStyle);

		let room1 = L.polygon([
			[25, 0],
			[40, 0],
			[40, 7],
			[25, 7]
		], footprintStyle);

		let room2 = L.polygon([
			[40, 0],
			[50.5, 0],
			[50.5, 15],
			[40, 15]
		], footprintStyle);

		let room3 = L.polygon([
			[12, 17],
			[31.5, 17],
			[31.5, 35],
			[12, 35]
		], footprintStyle);

		let room4 = L.polygon([
			[31.5, 26],
			[50.5, 26],
			[50.5, 40],
			[31.5, 40]
		], footprintStyle);

		let door1 = L.polygon([
			[32, 7],
			[34, 7],
			[34, 7],
			[32, 7]
		], {color: 'white'});

		let door2 = L.polygon([
			[44, 15],
			[46, 15],
			[46, 15],
			[44, 15]
		], {color: 'white'});

		let door3 = L.polygon([
			[21, 17],
			[23, 17],
			[23, 17],
			[21, 17]
		], {color: 'white'});

		let door4 = L.polygon([
			[40, 26],
			[42, 26],
			[42, 26],
			[40, 26]
		], {color: 'white'});

		let door5 = L.polygon([
			[0, 3],
			[0, 6],
			[0, 6],
			[0, 3]
		], {color: 'white'});

		const L1 = L.layerGroup([L1footprint, room1, room2, room3, room4, door1, door2, door3, door4, door5]);

		const map = L.map('map', {
			crs: L.CRS.Simple,
			minZoom: 0,
			layers: [L1]
		});

		map.setView([25.25, 9.5], 3);

		room1.on('mouseover', function() {
			this.setStyle({ fillColor: '#FFFB7EFF' });
		});
		room1.on('mouseout', function() {
			this.setStyle({ fillColor: 'white' });
		});

		room2.on('mouseover', function() {
			this.setStyle({ fillColor: '#FFFB7EFF' });
		});
		room2.on('mouseout', function() {
			this.setStyle({ fillColor: 'white' });
		});

		room3.on('mouseover', function() {
			this.setStyle({ fillColor: '#FFFB7EFF' });
		});
		room3.on('mouseout', function() {
			this.setStyle({ fillColor: 'white' });
		});

		room4.on('mouseover', function() {
			this.setStyle({ fillColor: '#FFFB7EFF' });
		});
		room4.on('mouseout', function() {
			this.setStyle({ fillColor: 'white' });
		});

		door5.on('mouseover', function() {
			this.setStyle({ color: '#FFFB7EFF' });
		});
		door5.on('mouseout', function() {
			this.setStyle({ color: 'white' });
		});

		const room1Label = L.marker([34, 2], { icon: L.divIcon({ className: 'room-label', html: 'Отдых' }) }).addTo(map);
		const room2Label = L.marker([45, 4], { icon: L.divIcon({ className: 'room-label', html: 'Руководство' }) }).addTo(map);
		const room3Label = L.marker([22, 20], { icon: L.divIcon({ className: 'room-label', html: 'Бухгалтерия' }) }).addTo(map);
		const room4Label = L.marker([41, 30], { icon: L.divIcon({ className: 'room-label', html: 'Переговорная' }) }).addTo(map);

	// unmount map function
	map.remove();
	}, []);

	return (
		<div>
			<style>
				{`
          html, body {
            height: 100%;
            margin: 0;
          }
          #map {
            width: 100%;
            height: 100%;
          }
          .imgclass {
            display: block;
            max-width: 400px;
            max-height: 400px;
            width: auto;
            height: auto;
          }
        `}
			</style>
			<div id='map'></div>
		</div>
	)
}

export default IndoorMap;
