---
carousels:
  - images:
    - image: images/connecting_apres.jpg
    - image: images/camera_testing.jpg
    - image: images/Yang_Testing.jpg
    - image: images/students.jpg
    - image: images/ravine.jpg
    - image: images/Tents.jpg
    - image: images/testing_with_Neo.jpg
    - image: images/radar_tripod_crop.jpg
    - image: images/summer_2025_testing.jpg
---

#  Computational Imaging in Photography and Environmental Radar

{% include carousel.html height="50" unit="%" duration="5" number="1" %}

{% include section.html %}
#  At the High precision Imaging, Signals N Radar HI SNR Lab, we develop computational imaging techniques that sense our environment at a scale, resolution, or quality that was previously impossible. 

## Research Areas

{% capture text %}

We are inventing radars, signal processing algorithms, and data assimilation techniques to measure the environment at scales previously unexplored. Our systems are deployed to improve geophysical understanding of the earth and to aid in resource management. We have developed systems that inform predictions for sea level rise and wildfire risk, as well as monitoring ecosystem and agricultural health.

{% 
  include button.html
  link="research"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/approachingradar.jpg"
  link="research"
  title="Environmental Radar"
  text=text
%}

{% capture text %}

We develop imaging and data assimilation algorithms not only for radar, but also for cameras. We develop equitable imaging algorithms that ensure beautiful, crisp images for all users. We also research 3D reconstruction methods, and compressive sensing. 


{%
  include button.html
  link="research"
  text="See our publications"
  icon="fa-solid fa-arrow-right"
  flip=true
  style="bare"
%}

{% endcapture %}

{%
  include feature.html
  image="images/camera_testing.jpg"
  link="research"
  title="Computational Photography"
  flip=true
  style="bare"
  text=text
%}
