/*
 * Copyright 2016-2018, EnMasse authors.
 * License: Apache License 2.0 (see the file LICENSE or http://apache.org/licenses/LICENSE-2.0.html).
 */
package io.enmasse.systemtest;

public interface AddressSpacePlan {
    String BROKERED = "brokered-single-broker";
    String STANDARD_SMALL = "standard-small";
    String STANDARD_MEDIUM = "standard-medium";
    String STANDARD_UNLIMITED = "standard-unlimited";
    String STANDARD_UNLIMITED_WITH_MQTT = "standard-unlimited-with-mqtt";
}
